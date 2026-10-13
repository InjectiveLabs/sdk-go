#!/usr/bin/env python3
"""Exercise credential scoping and private module proxy preparation offline."""

import base64
import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


SCRIPTS = Path(__file__).resolve().parents[1]


class MQBuildTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.scripts = self.root / "scripts"
        self.scripts.mkdir()
        for name in ("with-private-git.sh", "prepare-mq-go-modules.sh"):
            shutil.copy2(SCRIPTS / name, self.scripts / name)
        self.env = {
            key: value
            for key, value in os.environ.items()
            if not key.startswith("GIT_CONFIG_")
            and key not in ("GH_TOKEN", "GITHUB_ACTIONS", "GOPRIVATE", "GONOSUMDB")
        }

    def run_script(self, name, *args):
        return subprocess.run(
            [str(self.scripts / name), *args],
            env=self.env,
            text=True,
            capture_output=True,
            check=False,
        )

    def test_missing_ci_secret_fails_before_running_command(self):
        self.env["GITHUB_ACTIONS"] = "true"
        result = self.run_script("with-private-git.sh", "echo", "should-not-run")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("Missing GH_TOKEN secret", result.stdout)
        self.assertNotIn("should-not-run", result.stdout)

    def test_credential_is_command_scoped_and_token_variable_is_unset(self):
        self.env["GH_TOKEN"] = "fixture-token"
        result = self.run_script(
            "with-private-git.sh",
            "python3",
            "-c",
            "import json,os; print(json.dumps(dict(os.environ)))",
        )
        self.assertEqual(result.returncode, 0, result.stderr)
        child = json.loads(result.stdout)
        self.assertNotIn("GH_TOKEN", child)
        self.assertEqual(child["GIT_CONFIG_COUNT"], "1")
        self.assertEqual(
            child["GIT_CONFIG_KEY_0"],
            "http.https://github.com/InjectiveLabs/.extraheader",
        )
        expected = base64.b64encode(b"x-access-token:fixture-token").decode()
        self.assertEqual(child["GIT_CONFIG_VALUE_0"], "AUTHORIZATION: basic " + expected)
        self.assertEqual(child["GOPRIVATE"], "github.com/InjectiveLabs")
        self.assertNotIn("GIT_CONFIG_VALUE_0", self.env)

    def mock_go(self):
        self.cache = self.root / "go-cache"
        self.bin = self.root / "bin"
        self.bin.mkdir()
        self.log = self.root / "commands.jsonl"
        self.env.update(
            {
                "PATH": str(self.bin) + os.pathsep + self.env["PATH"],
                "FIXTURE_CACHE": str(self.cache),
                "FIXTURE_LOG": str(self.log),
            }
        )
        go = self.bin / "go"
        go.write_text(
            "#!/usr/bin/env python3\n"
            "import json,os,sys\n"
            "args = sys.argv[1:]\n"
            "with open(os.environ['FIXTURE_LOG'], 'a') as log:\n"
            "    log.write(json.dumps(args) + '\\n')\n"
            "if args == ['env', 'GOMODCACHE']:\n"
            "    print(os.environ['FIXTURE_CACHE'])\n"
            "elif args[:2] == ['list', '-m']:\n"
            "    print('github.com/InjectiveLabs/sdk-go')\n"
            "    print('github.com/InjectiveLabs/cosmos-sdk v0.50.15-inj.1')\n"
            "    print('github.com/InjectiveLabs/SomeModule v1.2.3')\n"
            "    print('example.com/public v2.0.0')\n"
            "elif args[:2] != ['mod', 'download']:\n"
            "    sys.exit('unexpected go command')\n"
        )
        go.chmod(0o755)
        for module, version in (
            ("cosmos-sdk", "v0.50.15-inj.1"),
            ("!some!module", "v1.2.3"),
        ):
            directory = self.cache / "cache/download/github.com/!injective!labs" / module / "@v"
            directory.mkdir(parents=True)
            for extension in ("info", "mod", "zip", "ziphash"):
                (directory / f"{version}.{extension}").write_text(extension)

    def test_proxy_contains_complete_private_versions_and_no_credentials(self):
        self.mock_go()
        self.env["GH_TOKEN"] = "fixture-token"
        result = self.run_script("prepare-mq-go-modules.sh")
        self.assertEqual(result.returncode, 0, result.stderr)
        proxy = self.root / ".modcache"
        files = [path for path in proxy.rglob("*") if path.is_file()]
        self.assertEqual(len(files), 8)
        self.assertTrue((proxy / "github.com/!injective!labs/!some!module/@v/v1.2.3.zip").is_file())
        for path in files:
            self.assertNotIn("fixture-token", path.read_text())
        commands = [json.loads(line) for line in self.log.read_text().splitlines()]
        self.assertIn(
            [
                "mod",
                "download",
                "github.com/InjectiveLabs/cosmos-sdk@v0.50.15-inj.1",
                "github.com/InjectiveLabs/SomeModule@v1.2.3",
            ],
            commands,
        )
        self.assertEqual(list(self.root.glob(".modcache.tmp.*")), [])

    def test_missing_artifact_preserves_previous_proxy_and_cleans_temporary_files(self):
        self.mock_go()
        missing = self.cache / "cache/download/github.com/!injective!labs/cosmos-sdk/@v/v0.50.15-inj.1.zip"
        missing.unlink()
        proxy = self.root / ".modcache"
        proxy.mkdir()
        (proxy / "previous").write_text("preserved")
        result = self.run_script("prepare-mq-go-modules.sh")
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("Missing downloaded module artifact", result.stderr)
        self.assertEqual((proxy / "previous").read_text(), "preserved")
        self.assertEqual(list(self.root.glob(".modcache.tmp.*")), [])


if __name__ == "__main__":
    unittest.main()

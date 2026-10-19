#!/usr/bin/env bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd -- "${script_dir}/.." && pwd)"
go_proxy_dir="${repo_root}/.modcache"
temp_go_proxy_dir="$(mktemp -d "${repo_root}/.modcache.tmp.XXXXXX")"
module_list_file="${temp_go_proxy_dir}/modules.txt"

cleanup() {
  rm -rf -- "${temp_go_proxy_dir}"
}
trap cleanup EXIT

cd -- "${repo_root}"
"${script_dir}/with-private-git.sh" go mod download
"${script_dir}/with-private-git.sh" go list -m \
  -f '{{if .Replace}}{{.Replace.Path}} {{.Replace.Version}}{{else}}{{.Path}} {{.Version}}{{end}}' all \
  > "${module_list_file}"

go_module_cache="$(go env GOMODCACHE)"
private_modules=()
while read -r module_path module_version; do
  [[ -n "${module_path}" && -n "${module_version}" ]] || continue
  [[ "${module_path}" == github.com/InjectiveLabs/* ]] || continue
  private_modules+=("${module_path}@${module_version}")
done < "${module_list_file}"

if [[ ${#private_modules[@]} -eq 0 ]]; then
  echo "No InjectiveLabs modules were found for the MQ Docker build." >&2
  exit 1
fi

# Restored Go caches can omit proxy zips; fetch each selected version explicitly.
"${script_dir}/with-private-git.sh" go mod download "${private_modules[@]}"

module_count=0
while read -r module_path module_version; do
  [[ -n "${module_path}" && -n "${module_version}" ]] || continue
  [[ "${module_path}" == github.com/InjectiveLabs/* ]] || continue

  # Escape every uppercase letter according to Go's module-cache convention.
  escaped_module_path="$(printf '%s' "${module_path}" | LC_ALL=C sed 's/[A-Z]/!&/g' | tr '[:upper:]' '[:lower:]')"
  source_version_dir="${go_module_cache}/cache/download/${escaped_module_path}/@v"
  target_version_dir="${temp_go_proxy_dir}/${escaped_module_path}/@v"
  mkdir -p -- "${target_version_dir}"

  for extension in info mod zip ziphash; do
    source_file="${source_version_dir}/${module_version}.${extension}"
    if [[ ! -f "${source_file}" ]]; then
      echo "Missing downloaded module artifact: ${source_file}" >&2
      exit 1
    fi
    cp -- "${source_file}" "${target_version_dir}/"
  done
  module_count=$((module_count + 1))
done < "${module_list_file}"

rm -- "${module_list_file}"
# Replace generated output only after every required module artifact is present.
rm -rf -- "${go_proxy_dir}"
mv -- "${temp_go_proxy_dir}" "${go_proxy_dir}"
trap - EXIT

echo "Prepared ${module_count} InjectiveLabs module versions in ${go_proxy_dir}."

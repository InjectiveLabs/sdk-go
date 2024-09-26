package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const (
	defaultOfacListURL      = "https://raw.githubusercontent.com/InjectiveLabs/injective-lists/master/wallets/ofac.json"
	defaultofacListFilename = "ofac.json"
)

var (
	ofacListFilename = defaultofacListFilename
)

type OfacChecker struct {
	ofacListPath string
	ofacList     map[string]bool
}

func NewOfacChecker() (*OfacChecker, error) {
	checker := &OfacChecker{
		ofacListPath: getOfacListPath(),
	}
	if _, err := os.Stat(checker.ofacListPath); os.IsNotExist(err) {
		if err := DownloadOfacList(); err != nil {
			return nil, err
		}
	}
	if err := checker.loadOfacList(); err != nil {
		return nil, err
	}
	return checker, nil
}

func getOfacListPath() string {
	return getFileAbsPath(path.Join("..", "metadata", ofacListFilename))
}

func DownloadOfacList() error {
	resp, err := http.Get(defaultOfacListURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download OFAC list, status code: %d", resp.StatusCode)
	}

	outFile, err := os.Create(getOfacListPath())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}
	_, err = outFile.WriteString("\n")
	if err != nil {
		return err
	}
	return nil
}

func (oc *OfacChecker) loadOfacList() error {
	file, err := os.ReadFile(oc.ofacListPath)
	if err != nil {
		return err
	}
	var list []string
	err = json.Unmarshal(file, &list)
	if err != nil {
		return err
	}
	oc.ofacList = make(map[string]bool)
	for _, item := range list {
		oc.ofacList[item] = true
	}
	return nil
}

func (oc *OfacChecker) IsBlacklisted(address string) bool {
	return oc.ofacList[address]
}

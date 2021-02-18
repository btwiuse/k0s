package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func expandHomePath(fp string) (string, error) {
	if fp[0] != '~' {
		return fp, nil
	}
	hd, err := os.UserHomeDir()
	if err != nil {
		return fp, err
	}
	fp = filepath.Join(hd, fp[1:])
	return fp, nil
}

func ReadFileBytes(fp string) ([]byte, error) {
	var err error
	fp, err = expandHomePath(fp)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(fp)
}

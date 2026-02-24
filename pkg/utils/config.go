package utils

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// FileExists checks if a file exists and is not a directory.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// ProbeConfigFiles searches for the first existing config file
// from the provided list of paths.
func ProbeConfigFiles(paths []string) string {
	for _, path := range paths {
		if FileExists(path) {
			return path
		}
	}
	return ""
}

// LoadYAMLConfig loads a YAML configuration file into the target struct.
// If the file is empty or does not exist, it returns without error.
// Returns error only if the file exists but cannot be read or parsed.
func LoadYAMLConfig(file string, target interface{}) error {
	if file == "" {
		return nil
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	err = dec.Decode(target)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

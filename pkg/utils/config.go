package utils

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// FileExists checks if a path exists in the filesystem.
// It returns true for both files and directories.
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	// Return true if path exists, but false if it's a directory
	// (config files should be regular files, not directories)
	return err == nil && !info.IsDir()
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

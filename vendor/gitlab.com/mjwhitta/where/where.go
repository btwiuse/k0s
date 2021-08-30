package where

import (
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/mjwhitta/pathname"
)

// Is will return the full path to the specified cmd if it exists in
// the defined PATH env var.
func Is(cmd string) string {
	var cached interface{}
	var dirs []string
	var exts []string
	var fullpath string
	var hasKey bool

	if cmd == "" {
		return ""
	}

	// Return cached value if it exists
	if cached, hasKey = cache.Get(cmd); hasKey {
		return cached.(string)
	}

	// Get all PATH directories
	dirs = strings.Split(
		os.Getenv("PATH"),
		string(os.PathListSeparator),
	)

	// Get all valid extensions
	exts = strings.Split(os.Getenv("PATHEXT"), ";")
	if len(exts) == 0 {
		exts = append(exts, "")
	}

	// Find, cache, and return full path
	for _, dir := range dirs {
		for _, ext := range exts {
			fullpath = filepath.Join(dir, cmd+ext)

			if pathname.DoesExist(fullpath) {
				cache.Put(cmd, pathname.ExpandPath(fullpath))
				cached, _ = cache.Get(cmd)
				return cached.(string)
			}
		}
	}

	// Otherwise return empty string
	return ""
}

package pathname

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Basename wraps filepath.Base(path str).
func Basename(path string) string {
	return filepath.Base(ExpandPath(path))
}

// Dirname wraps filepath.Dir(path str).
func Dirname(path string) string {
	return filepath.Dir(ExpandPath(path))
}

// DoesExist returns true if the specified path exists on disk, false
// otherwise.
func DoesExist(path string) bool {
	if _, err := os.Stat(ExpandPath(path)); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// ExpandPath will expand the specified path accounting for ~ or ~user
// shortcuts as well as ENV vars.
func ExpandPath(path string) string {
	var e error
	var sep int
	var usr *user.User

	// Fix separators
	path = filepath.Clean(path)

	// Expand ENV vars
	path = os.ExpandEnv(path)

	// Expand ~
	if strings.HasPrefix(path, "~") {
		sep = strings.Index(path, string(filepath.Separator))

		switch sep {
		case -1:
			// If just ~
			if path == "~" {
				if usr, e = user.Current(); e != nil {
					return path
				}
			} else {
				// If ~user shortcut
				if usr, e = user.Lookup(path[1:]); e != nil {
					return path
				}
			}

			path = ""
		case 1:
			// If path starting with ~/
			if usr, e = user.Current(); e != nil {
				return path
			}

			path = path[2:]
		default:
			// If ~user/ shortcut
			if usr, e = user.Lookup(path[1:sep]); e != nil {
				return path
			}

			path = path[sep+1:]
		}

		return filepath.Join(usr.HomeDir, path)
	}

	// Otherwise just return path
	return path
}

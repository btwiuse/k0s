package pathname

import (
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"gitlab.com/mjwhitta/errors"
)

// Basename wraps filepath.Base(path str).
func Basename(path string) string {
	return filepath.Base(ExpandPath(path))
}

// Dirname wraps filepath.Dir(path str).
func Dirname(path string) string {
	return filepath.Dir(ExpandPath(path))
}

// DoesExist returns whether or not the file exists on disk. An error
// is returned, if the path is not accessible.
func DoesExist(path string) (bool, error) {
	var e error

	if _, e = os.Stat(ExpandPath(path)); e == nil {
		return true, nil
	} else if os.IsNotExist(e) {
		return false, nil
	}

	e = e.(*fs.PathError).Err
	e = errors.Newf("path %s not accesssible: %w", path, e)

	return false, e
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

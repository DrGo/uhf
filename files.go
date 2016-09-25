package uhf

import "os"

// IsDir accepts a string (file path) and returns
// a boolean which indicates if the path is
// a valid directory.
func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

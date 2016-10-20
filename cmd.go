package uhf

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

var ErrNoInput = errors.New("no input provided")

// IsInteractive returns a boolean
// indicating whether os.Stdin is
// a user session. If true, then the
// user is running the program interactively.
//
// If false, data is being piped into the program.
func IsInteractive() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return stat.Mode()&os.ModeCharDevice != 0, nil
}

// BinDir returns the absolute path of the
// directory containing the binary currently running.
func BinDir() (string, error) {
	return filepath.Abs(os.Args[0])
}

// FileExists returns a boolean indicating whether
// a file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// InputReader returns a reader from the
// application's input, whether this is
// stdin or read from files on the command line.
func InputReader() (io.Reader, error) {
	if b, err := IsInteractive(); err == nil {
		if !b {
			return os.Stdin, nil
		}
	}

	readers := make([]io.Reader, 0, 1)
	for _, filename := range os.Args[1:] {
		if FileExists(filename) {
			f, err := os.Open(filename)
			if err != nil {
				continue
			}
			readers = append(readers, f)
		}
	}
	if len(readers) == 0 {
		return io.MultiReader(readers...), ErrNoInput
	}
	return io.MultiReader(readers...), nil
}

package uhf

import "os"

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

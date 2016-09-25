package uhf

import "testing"

func TestIsInteractive(t *testing.T) {
	if result, err := IsInteractive(); err != nil {
		t.Fatalf("failed to test IsInteractive: %s\n", err)
	} else if !result {
		t.Fatalf("Expected to not be interactive during test run.\n")
	}
}

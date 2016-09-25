package uhf

import "testing"

func TestIsInteractive(t *testing.T) {
	if result, err := IsInteractive(); err != nil {
		t.Fatalf("failed to test IsInteractive: %s\n", err)
	} else if !result {
		t.Fatalf("Expected to not be interactive during test run.\n")
	}
}

func TestBinDir(t *testing.T) {
	path, err := BinDir()
	if err != nil {
		t.Fatalf("failed to test BinDir: %s\n", err)
	}
	if path == "" {
		t.Fatal("Expected path from BinDir, got nothing.")
	}
}

func TestFileExists(t *testing.T) {
	wanted := []struct {
		path   string
		exists bool
	}{
		{"README.md", true},
		{"fake.txt", false},
	}

	for _, item := range wanted {
		if FileExists(item.path) != item.exists {
			t.Fatalf("expected %t for %q.\n", item.exists, item.path)
		}
	}
}

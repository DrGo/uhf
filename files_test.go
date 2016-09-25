package uhf

import "testing"

func TestIsDir(t *testing.T) {
	wanted := []struct {
		path        string
		isDir       bool
		shouldError bool
	}{
		{".", true, false},
		{"./README.md", false, false},
		{"/root/.ssh/authorized_keys", false, true},
	}

	for _, item := range wanted {
		result, err := IsDir(item.path)
		if err != nil {
			if !item.shouldError {
				t.Fatalf("failed to test IsDir: %s\n", err)
			}
			continue
		}

		if result != item.isDir {
			t.Fatalf("Expected %q isDir %t, got %t\n", item.path, item.isDir, result)
		}
	}
}

package uhf

import "testing"

func TestCompress(t *testing.T) {
	long := "                                                          "
	result, err := Compress([]byte(long))
	if err != nil {
		t.Fatalf("compress failed: %s\n", err)
	}

	/*
		short := string(result)
			if len(long) > len(short) {
				t.Fatalf("long: %d, short %d %q, result %#v\n", len(long), len(short), short, result)
			}
	*/

	inflated, err := Decompress(result)
	if err != nil {
		t.Fatalf("decompress failed: %s\n", err)
	}
	expanded := string(inflated)

	if expanded != long {
		t.Fatalf("expanded: %q, long %q", expanded, long)
	}

}

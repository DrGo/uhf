package uhf

import (
	"io"
	"testing"
)

func TestCSVSliceChan(t *testing.T) {
	numRecs := 3
	var recs [][]string
	reader := CSVToSliceChan("test_data/sample.csv")
	for rec := range reader.C {
		recs = append(recs, rec)
	}
	if len(recs) != numRecs {
		t.Fatalf("Expected %d records, got %d\n", numRecs, len(recs))
	}

	if reader.Error() != io.EOF {
		t.Fatalf("Expected EOF, got %s\n", reader.Error())
	}
}

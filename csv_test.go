package uhf

import (
	"io"
	"testing"
)

var numRecs = 3

func TestCSVSliceChan(t *testing.T) {
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

func TestCSVToSlice(t *testing.T) {
	recs, err := CSVToSlice("test_data/sample.csv")
	if len(recs) != numRecs {
		t.Fatalf("Expected %d records, got %d\n", numRecs, len(recs))
	}

	if err != io.EOF {
		t.Fatalf("Expected EOF, got %s\n", err)
	}
}

func TestCSVMapChan(t *testing.T) {
	var recs []map[string]string
	reader := CSVToMapChan("test_data/sample.csv")
	for rec := range reader.C {
		recs = append(recs, rec)
	}
	if len(recs) != numRecs-1 {
		t.Fatalf("Expected %d records, got %d\n", numRecs, len(recs))
	}

	if reader.Error() != io.EOF {
		t.Fatalf("Expected EOF, got %s\n", reader.Error())
	}

	if recs[0]["Title"] != "Word Crimes" {
		t.Fatalf("Expected \"Word Crimes,\", got %#v\n", recs[0])
	}
}

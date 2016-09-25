package uhf

import (
	"io/ioutil"
	"os"
	"testing"
)

type person struct {
	First  string
	Middle string
	Last   string
	Nick   string
}

func TestLoadJSON(t *testing.T) {
	al := &person{}
	err := LoadJSON("test_data/sample.json", al)
	if err != nil {
		t.Fatalf("Error %s in TestLoadJSON\n", err)
	}

	if al.Nick != "Weird Al" {
		t.Fatalf("Wrong nickname %q\n", al.Nick)
	}
}

func TestSaveJSON(t *testing.T) {
	al := &person{}
	err := LoadJSON("test_data/sample.json", al)
	if err != nil {
		t.Fatalf("Error %s in loading JSON in TestSaveJSON\n", err)
	}
	tempfile, err := tempFilename()
	if err != nil {
		t.Fatalf("Error %s creating temp file in TestSaveJSON\n", err)
	}

	defer func() {
		_ = os.Remove(tempfile)
	}()

	err = SaveJSON(tempfile, al)
	if err != nil {
		t.Fatalf("Error %s saving clone file in TestSaveJSON\n", err)
	}

	clone := &person{}
	err = LoadJSON(tempfile, clone)
	if err != nil {
		t.Fatalf("Error %s in loading clone JSON in TestSaveJSON\n", err)
	}

	if al.Nick != clone.Nick {
		t.Fatalf("Nicknames didn't match, wanted %q, got %q\n", al.Nick, clone.Nick)
	}

}

func tempFilename() (string, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	f.Close()
	return f.Name(), nil
}

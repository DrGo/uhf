package uhf

import (
	"encoding/csv"
	"os"
	"sync"
)

// CSVChan provides a channel which will
// pass each record from a CSV file. When
// the channel is closed, the Error() function
// will return the error provided by the encoding/csv package.
type CSVChan struct {
	C    chan []string
	err  error
	lock sync.RWMutex
}

// Error returns the encoding/csv.Reader read error.
// If all goes well, this should be io.EOF.
func (c *CSVChan) Error() error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.err
}

func (c *CSVChan) setError(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.err = err
}

// CSVToSliceChan is a convenience for reading
// a CSV file into slices of strings without
// loading the entire CSV file into memory.
func CSVToSliceChan(filename string) *CSVChan {
	out := newCSVChan()
	f, err := os.Open(filename)
	if err != nil {
		out.setError(err)
		return out
	}

	go func() {
		defer f.Close()
		defer close(out.C)
		r := csv.NewReader(f)
		for {
			rec, err := r.Read()
			if err != nil {
				out.setError(err)
				break
			}
			out.C <- rec
		}
	}()

	return out
}

func newCSVChan() *CSVChan {
	return &CSVChan{C: make(chan []string)}
}

// CSVToSlice is a convenience function for slurping
// a CSV file into a [][]string. It loads the whole file
// at once, so this shouldn't be used with huge files.
func CSVToSlice(filename string) ([][]string, error) {
	var recs [][]string
	reader := CSVToSliceChan(filename)
	for rec := range reader.C {
		recs = append(recs, rec)
	}
	return recs, reader.Error()
}

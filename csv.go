package uhf

import (
	"encoding/csv"
	"errors"
	"os"
	"sync"
)

var ErrColMismatch = errors.New("number of columns in CSV do not match header")

// CSVSliceChan provides a channel which will
// pass each record from a CSV file. When
// the channel is closed, the Error() function
// will return the error provided by the encoding/csv package.
type CSVSliceChan struct {
	C    chan []string
	err  error
	lock sync.RWMutex
}

// CSVMapChan contains a channel to return maps
// built from a CSV file, using the header row
// as the keys. The Error method will return the
// error returned by the encoding/csv package when
// it finishes reading.
type CSVMapChan struct {
	C    chan map[string]string
	err  error
	lock sync.RWMutex
}

func (c *CSVMapChan) setError(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.err = err
}

// Error returns the encoding/csv.Reader read error.
// If all goes well, this should be io.EOF.
func (c *CSVSliceChan) Error() error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.err
}

func (c *CSVSliceChan) setError(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.err = err
}

// Error returns the encoding/csv.Reader read error.
// If all goes well, this should be io.EOF.
func (c *CSVMapChan) Error() error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.err
}

// CSVToSliceChan is a convenience for reading
// a CSV file into slices of strings without
// loading the entire CSV file into memory.
func CSVToSliceChan(filename string) *CSVSliceChan {
	out := &CSVSliceChan{C: make(chan []string)}
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

// CSVToMapChan is a convenience for reading
// a CSV file into maps.
func CSVToMapChan(filename string) *CSVMapChan {
	out := &CSVMapChan{C: make(chan map[string]string)}
	f, err := os.Open(filename)
	if err != nil {
		out.setError(err)
		return out
	}

	go func() {
		defer f.Close()
		defer close(out.C)
		r := csv.NewReader(f)
		headers, err := r.Read()
		if err != nil {
			out.setError(err)
			return
		}
		numCols := len(headers)
		for {
			m := make(map[string]string)
			rec, err := r.Read()
			if err != nil {
				out.setError(err)
				break
			}
			if len(rec) != numCols {
				out.setError(ErrColMismatch)
			}
			for i, h := range headers {
				m[h] = rec[i]
			}
			out.C <- m
		}
	}()

	return out
}

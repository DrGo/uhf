package uhf

import (
	"io/ioutil"
	"net/http"
)

// DownloadFile accepts a URL and filename and
// downloads the contentns of the URL to the
// filename if possible.
func DownloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

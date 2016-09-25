package uhf

import (
	"encoding/json"
	"io/ioutil"
)

// SaveJSON writens an inteface{} to a file in JSON format.
func SaveJSON(filename string, val interface{}) error {
	j, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, j, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadJSON loads an interface{} from a JSON file.
func LoadJSON(filename string, val interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, val)
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"soundcast/api/interfaces/data"
)

// JSONData is a structure that simplify the data access from a JSON file
type JSONData struct {
	path       string
	loadedData []data.DbElement
}

// LoadFile function allows to load a JSON file from a given path
func (da *JSONData) LoadFile(path string) (err error) {
	da.path = path
	jsonFile, err := os.Open(da.path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteValue, &da.loadedData)
	if err != nil {
		return err
	}
	return nil
}

// All function returns a slice of DbElement that matches the function passed as parameter
func (da *JSONData) All(f data.Matches) (results []data.DbElement) {
	for _, element := range da.loadedData {
		isMatch := f(element)
		if isMatch {
			results = append(results, element)
		}
	}
	return results
}

// First function returns the first occurence of DbElement that matches the function passed as parameter
func (da *JSONData) First(f data.Matches) data.DbElement {
	for _, element := range da.loadedData {
		isMatch := f(element)
		if isMatch {
			return element
		}
	}
	return nil
}

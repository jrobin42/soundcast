package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Represent an element of the database
type DbElement map[string]interface{}

// Structure that simplify the data access
type DataAccessor struct {
	path       string
	loadedData []DbElement
}
// This function allows to load a JSON file from a given path
func (da *DataAccessor) LoadJSON(path string) (err error) {
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

// Should return true if the given element matches search criterias
type Matches func (DbElement) bool

// This function returns a slice of DbElement that matches the function passed as parameter
func (da *DataAccessor) FindAll(f Matches) (results []DbElement) {
	for _, element := range da.loadedData {
		isMatch := f(element)
		if isMatch {
			results = append(results, element)
		}
	}
	return results
}
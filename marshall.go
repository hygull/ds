package ds

import (
	"encoding/json"
)

//A function that marshalls the parameter data and returns the string representation
func Marshall(data interface{}) (string, error) {
	dataSlice, err := json.MarshalIndent(data, "", "\t")
	return string(dataSlice), err
}

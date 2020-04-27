package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// PrintAndLog writes to stdout and to a logger.
func PrintAndLog(message string) {
	log.Println(message)
	fmt.Println(message)
}

func LogAndPanic(err error) {
	PrintAndLog(err.Error())
	panic(err)
}

// ReadJSON reads a json file, and unmashals it.
// Very useful for template deployments.
func ReadJSON(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read template file: %v\n", err)
	}
	contents := make(map[string]interface{})
	if err := json.Unmarshal(data, &contents); err != nil {
		return nil, err
	}
	return &contents, nil
}

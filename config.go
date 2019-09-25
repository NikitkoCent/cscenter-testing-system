package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)


type Test struct {
	Name string `json:"name"`
	ExecParams []string `json:"params"`
	ReferencePath *string `json:"reference"`
	TestedFilePath *string `json:"tested_file"`
	ExitCodes []int `json:"exitCodes"`
}

type TestsConfig struct {
	Executable *string `json:"executable"`
	Tests []Test `json:"tests"`
}


func ReadTestsConfig(path string) (TestsConfig, error) {
	result := TestsConfig{}

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(bytes, &result); err != nil {
		return result, err
	}

	createNotFoundFieldError := func (fieldName string) error {
		return errors.New(fmt.Sprintf("'%s': field '%s' was not found.", path, fieldName))
	}

	createTestNotFoundFieldError := func (fieldName string, testName string) error {
		return errors.New(fmt.Sprintf("'%s::%s': field '%s' was not found.", path, testName, fieldName))
	}

	if result.Executable == nil {
		return result, createNotFoundFieldError("executable")
	}
	for i := range result.Tests {
		if len(result.Tests[i].Name) == 0 {
			result.Tests[i].Name = fmt.Sprintf("Unnamed~%d", i)
		}

		if result.Tests[i].ReferencePath == nil {
			return result, createTestNotFoundFieldError("reference", result.Tests[i].Name)
		}

		if len(result.Tests[i].ExitCodes) == 0 {
			result.Tests[i].ExitCodes = []int{0}
		}
	}

	return result, nil
}

package main

import (
	"io/ioutil"
	"os"
)

type FileProcessor struct {
}

func (f FileProcessor) Process(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return NewFileCouldNotBeOpenedError(err.Error())
	}

	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	/* Add a check for the file being empty and return a FileWasEmptyError */

	return nil
}

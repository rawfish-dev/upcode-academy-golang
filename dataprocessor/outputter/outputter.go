package outputter

import (
	"log"
	"os"
)

type Outputter struct {
	logger   *log.Logger
	FileName string
}

func NewOutputter(l *log.Logger, f string) *Outputter {
	return &Outputter{
		logger:   l,
		FileName: f,
	}
}

func (o *Outputter) Output(sortedData []string) error {
	file, err := os.OpenFile(o.FileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add a header to the file by inserting a string here

	for idx := range sortedData {
		file.WriteString(sortedData[idx] + "\n")
	}

	return nil
}

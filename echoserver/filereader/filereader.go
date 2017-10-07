package filereader

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	/* An import is missing here */
	"upcode-academy-golang/echoserver/common"
)

var _ interfaces.DataProvider = new(FileReader)

type FileReader struct {
	/* A field name is missing here */
	Data string
}

func NewFileReader(fileName string) *FileReader {
	return &FileReader{
		FileName: fileName,
	}
}

/* An interface method is missing here */

func (f *FileReader) ReadAll() []string {
	file, err := os.Open(f.FileName)
	if err != nil {
		log.Fatalf("unable to open file %s due to %v", f.FileName, err)
	}

	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file contents due to %v", err)
	}

	/* A string operation for removing leading and trailing spaces is missing here */

	/* A string operation required to divide the content in a slice is missing here */

	return nil
}

package filereader

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"upcode-academy-golang/echoserver_full/common"
	"upcode-academy-golang/echoserver_full/interfaces"
)

var _ interfaces.DataProvider = new(FileReader)

type FileReader struct {
	FileName string
	Data     string
}

func NewFileReader(fileName string) *FileReader {
	return &FileReader{
		FileName: fileName,
	}
}

func (f *FileReader) BuildMessagePad() string {
	return common.BuildMessagePad("File Reader")
}

func (f *FileReader) ReadAll() []string {
	file, err := os.Open(f.FileName)
	if err != nil {
		log.Fatalf("unable to open file %s due to %v", f.FileName, err)
	}

	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file contents due to %v", err)
	}

	trimmedContent := strings.TrimSpace(string(contentBytes))

	return strings.Split(trimmedContent, "\n")
}

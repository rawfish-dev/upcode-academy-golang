package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"upcode-academy-golang/echoserver_full/common"
	"upcode-academy-golang/echoserver_full/filereader"
	"upcode-academy-golang/echoserver_full/interfaces"
	"upcode-academy-golang/echoserver_full/memory"
)

var (
	httpClient *http.Client
)

func init() {
	httpClient = http.DefaultClient
}

func main() {
	dataProviders := []interfaces.DataProvider{
		memory.NewMemoryBuffer(),
		filereader.NewFileReader("filereader/data.txt"),
	}

	log.Println("Beginning to send data...")

	sendAll(dataProviders)

	log.Println("All data sent!")
}

func sendAll(dataProviders []interfaces.DataProvider) {
	for _, dataProvider := range dataProviders {
		sendDataWithProvider(dataProvider)
	}
}

func sendDataWithProvider(dataProvider interfaces.DataProvider) {
	data := dataProvider.ReadAll()

	data = append([]string{dataProvider.BuildMessagePad()}, data...)

	for _, value := range data {
		constructedMessage := common.Message{
			Value: value,
		}

		dataBytes, err := json.Marshal(constructedMessage)
		if err != nil {
			log.Fatalf("unable to marshal data due to %v", err)
		}

		_, err = httpClient.Post("http://localhost:8080/messages", "application/json", bytes.NewReader(dataBytes))
		if err != nil {
			log.Fatalf("unable to complete http request due to %v", err)
		}
	}
}

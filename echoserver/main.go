package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"upcode-academy-golang/echoserver/common"
	"upcode-academy-golang/echoserver/filereader"
	"upcode-academy-golang/echoserver/interfaces"
	"upcode-academy-golang/echoserver/memory"
)

var (
	httpClient *http.Client
)

func init() {
	log.Println("Initialising http client!")
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

	messagePad := dataProvider.BuildMessagePad()

	data = append([]string{messagePad}, data...)

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

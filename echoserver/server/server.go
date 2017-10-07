package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"upcode-academy-golang/echoserver/common"

	"github.com/gorilla/mux"
)

var (
	echoedValues []string
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/messages", messageHandler).Methods("POST")
	router.HandleFunc("/echo", echoHandler).Methods("GET")

	log.Println("Starting up the server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received content with length %d", r.ContentLength)

	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("could not read request body due to %v", err)
	}

	var message common.Message
	err = json.Unmarshal(requestBytes, &message)
	if err != nil {
		log.Fatalf("could not unmarshal request bytes due to %v", err)
	}

	fullMessage := fmt.Sprintf("%s: %s", time.Now().UTC(), message.Value)

	echoedValues = append(echoedValues, fullMessage)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if len(echoedValues) == 0 {
		w.Write([]byte("Nothing to echo yet..."))
	} else {
		for _, value := range echoedValues {
			w.Write([]byte(value + "\n"))
		}
	}
}

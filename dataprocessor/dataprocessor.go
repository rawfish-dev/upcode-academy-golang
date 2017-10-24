package main

import (
	"log"
	"os"
	"time"

	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/outputter"
	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/processor"
	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/reader"
)

func main() {
	logFile, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	startTime := time.Now()

	reader := reader.NewReader(logger)
	countryData, stateData, cityData := reader.ReadAll()

	processor := processor.NewProcessor(logger)
	processedData := processor.Process(countryData, stateData, cityData)

	outputter := outputter.NewOutputter(logger, "processed_data.csv")
	err = outputter.Output(processedData)
	if err != nil {
		logger.Printf("unable to complete outputting due to %s", err.Error())
	}

	logger.Printf("completed data processing in %0.2f seconds", time.Since(startTime).Seconds())
}

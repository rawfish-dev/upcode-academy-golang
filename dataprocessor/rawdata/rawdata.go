package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/common"
)

const folderName = "Countries-States-Cities-database"

func main() {
	generateCountryCSV()
	generateStateCSV()
	generateCityCSV()
}

func generateCountryCSV() {
	countryFile, err := os.Open(filepath.Join(folderName, "countries.json"))
	if err != nil {
		panic(err)
	}

	countryFileContents, err := ioutil.ReadAll(countryFile)
	if err != nil {
		panic(err)
	}

	var countries common.CountryWrapper
	err = json.Unmarshal(countryFileContents, &countries)
	if err != nil {
		panic(err)
	}

	processedCountryFile, err := os.OpenFile("processed_countries.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer processedCountryFile.Close()

	headers := "id,name,code"
	processedCountryFile.WriteString(headers + "\n")

	rowData := make([]string, 0, len(countries.Countries))
	for idx := range countries.Countries {
		rowData = append(rowData, countries.Countries[idx].ID+","+countries.Countries[idx].Name+","+countries.Countries[idx].CountryCode+"\n")
	}

	jumbledRowData := make([]string, len(rowData))
	perm := rand.Perm(len(rowData))
	for i, v := range perm {
		jumbledRowData[v] = rowData[i]
	}

	for idx := range jumbledRowData {
		processedCountryFile.WriteString(jumbledRowData[idx])
	}
}

func generateStateCSV() {
	stateFile, err := os.Open(filepath.Join(folderName, "states.json"))
	if err != nil {
		panic(err)
	}

	stateFileContents, err := ioutil.ReadAll(stateFile)
	if err != nil {
		panic(err)
	}

	var states common.StateWrapper
	err = json.Unmarshal(stateFileContents, &states)
	if err != nil {
		panic(err)
	}

	processedStateFile, err := os.OpenFile("processed_states.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer processedStateFile.Close()

	headers := "id,name,country_id"
	processedStateFile.WriteString(headers + "\n")

	rowData := make([]string, 0, len(states.States))
	for idx := range states.States {
		rowData = append(rowData, states.States[idx].ID+","+states.States[idx].Name+","+states.States[idx].CountryID+"\n")
	}

	jumbledRowData := make([]string, len(rowData))
	perm := rand.Perm(len(rowData))
	for i, v := range perm {
		jumbledRowData[v] = rowData[i]
	}

	for idx := range jumbledRowData {
		processedStateFile.WriteString(jumbledRowData[idx])
	}
}

func generateCityCSV() {
	cityFile, err := os.Open(filepath.Join(folderName, "cities.json"))
	if err != nil {
		panic(err)
	}

	cityFileContents, err := ioutil.ReadAll(cityFile)
	if err != nil {
		panic(err)
	}

	var cities common.CityWrapper
	err = json.Unmarshal(cityFileContents, &cities)
	if err != nil {
		panic(err)
	}

	processedCityFile, err := os.OpenFile("processed_cities.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer processedCityFile.Close()

	headers := "id,name,state_id"
	processedCityFile.WriteString(headers + "\n")

	rowData := make([]string, 0, len(cities.Cities))
	for idx := range cities.Cities {
		rowData = append(rowData, cities.Cities[idx].ID+","+cities.Cities[idx].Name+","+cities.Cities[idx].StateID+"\n")
	}

	jumbledRowData := make([]string, len(rowData))
	perm := rand.Perm(len(rowData))
	for i, v := range perm {
		jumbledRowData[v] = rowData[i]
	}

	for idx := range jumbledRowData {
		processedCityFile.WriteString(jumbledRowData[idx])
	}
}

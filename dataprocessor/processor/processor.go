package processor

import (
	"log"
	"sort"
	"strconv"

	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/common"
)

type Processor struct {
	logger *log.Logger
}

func NewProcessor(l *log.Logger) *Processor {
	return &Processor{
		logger: l,
	}
}

func (p *Processor) Process(rawCountries []common.Country, rawStates []common.State, rawCities []common.City) []string {
	// Country ID -> Country
	mappedCountries := ProcessCountries(rawCountries)

	// CountryID -> []State
	mappedStates := ProcessStates(rawStates)

	// StateID -> []City
	mappedCities := ProcessCities(rawCities)

	// Why?
	processedData := make([]common.RowData, 0, len(rawCountries))
	for countryID, country := range mappedCountries {
		for _, state := range mappedStates[countryID] {
			for _, city := range mappedCities[state.ID] {
				rowData := common.RowData{
					CountryName: country.Name,
					StateName:   state.Name,
					CityName:    city.Name,
				}
				processedData = append(processedData, rowData)
			}
		}
	}

	sortedRowData := SortRowData(processedData)

	PrependID(sortedRowData)

	return sortedRowData
}

func SortRowData(rowData []common.RowData) []string {
	csvRows := make([]string, 0, len(rowData))
	for idx := range rowData {
		csvRow := rowData[idx].CountryName + "," + rowData[idx].StateName + "," + rowData[idx].CityName
		csvRows = append(csvRows, csvRow)
	}

	sort.Strings(csvRows)

	// Sorted
	return csvRows
}

func PrependID(sortedData []string) {
	for idx := range sortedData {
		sortedData[idx] = strconv.Itoa(idx+1) + "," + sortedData[idx]
	}
}

func ProcessCountries(countries []common.Country) map[string]common.Country {
	mappedCountries := make(map[string]common.Country)
	for idx := range countries {
		mappedCountries[countries[idx].ID] = countries[idx]
	}

	return mappedCountries
}

func ProcessStates(states []common.State) map[string][]common.State {
	mappedStates := make(map[string][]common.State)
	for idx := range states {
		mappedStates[states[idx].CountryID] = append(mappedStates[states[idx].CountryID], states[idx])
	}

	return mappedStates
}

func ProcessCities(cities []common.City) map[string][]common.City {
	mappedCities := make(map[string][]common.City)
	for idx := range cities {
		mappedCities[cities[idx].StateID] = append(mappedCities[cities[idx].StateID], cities[idx])
	}

	return mappedCities
}

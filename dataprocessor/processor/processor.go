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
	/*

		Complete the entire implementation for Process

	*/
}

func SortRowData(rowData []common.RowData) []string {
	/*

		Complete the entire implementation for SortRowData

	*/
}

func PrependID(sortedData []string) {
	/*

		Complete the entire implementation for PrependID

	*/
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
	/*

		Complete the entire implementation for ProcessCities

	*/
}

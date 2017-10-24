package reader

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	"github.com/rawfish-dev/upcode-academy-golang/dataprocessor/common"
)

const (
	folderName = "./rawdata"
)

type Reader struct {
	logger *log.Logger
}

func NewReader(l *log.Logger) *Reader {
	return &Reader{
		logger: l,
	}
}

// Refactor to have a single method to do the file reading that returns an error if something fails

func (r *Reader) ReadAll() ([]common.Country, []common.State, []common.City) {
	return r.readCountries(), r.readStates(), r.readCities()
}

func (r *Reader) readCountries() []common.Country {
	file, err := os.Open(filepath.Join(folderName, "processed_countries.csv"))
	if err != nil {
		r.logger.Fatal(err)
	}

	lines, err := csv.NewReader(file).ReadAll()

	countries := make([]common.Country, 0, len(lines))
	for idx, line := range lines {
		// Skip header line
		if idx == 0 {
			continue
		}

		if len(line) != 3 {
			r.logger.Printf("line %s did not have the right number of values, skipping...", line)
		}

		country := common.Country{
			ID:          line[0],
			Name:        line[1],
			CountryCode: line[2],
		}

		countries = append(countries, country)
	}

	return countries
}

func (r *Reader) readStates() []common.State {
	file, err := os.Open(filepath.Join(folderName, "processed_states.csv"))
	if err != nil {
		r.logger.Fatal(err)
	}

	lines, err := csv.NewReader(file).ReadAll()

	states := make([]common.State, 0, len(lines))
	for idx, line := range lines {
		// Skip header line
		if idx == 0 {
			continue
		}

		if len(line) != 3 {
			r.logger.Printf("line %s did not have the right number of values, skipping...", line)
			continue
		}

		state := common.State{
			ID:        line[0],
			Name:      line[1],
			CountryID: line[2],
		}

		states = append(states, state)
	}

	return states
}

func (r *Reader) readCities() []common.City {
	file, err := os.Open(filepath.Join(folderName, "processed_cities.csv"))
	if err != nil {
		r.logger.Fatal(err)
	}

	lines, err := csv.NewReader(file).ReadAll()

	cities := make([]common.City, 0, len(lines))
	for idx, line := range lines {
		// Skip header line
		if idx == 0 {
			continue
		}

		if len(line) != 3 {
			r.logger.Printf("line %s did not have the right number of values, skipping...", line)
		}

		city := common.City{
			ID:      line[0],
			Name:    line[1],
			StateID: line[2],
		}

		cities = append(cities, city)
	}

	return cities
}

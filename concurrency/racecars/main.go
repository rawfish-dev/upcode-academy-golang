package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	totalLaps     = 10
	lapDistanceKM = 50
)

type Racecar struct {
	Brand       string
	KMPerSecond float64
	Completed   bool
}

type RaceResults struct {
	Positions []Result
	StartTime time.Time
	EndTime   time.Time
}

type Result struct {
	Racecar
	TimeCompleted time.Time
}

type LiveDriverReport struct {
	CurrentLap      float64
	CurrentDistance float64
}

func main() {
	rand.Seed(time.Now().Unix())

	// Prepare the cars
	toyota := Racecar{
		Brand:       "Toyota",
		KMPerSecond: 100,
	}
	subaru := Racecar{
		Brand:       "Subaru",
		KMPerSecond: 115,
	}
	porsche := Racecar{
		Brand:       "Porsche",
		KMPerSecond: 125,
	}
	ferrari := Racecar{
		Brand:       "Ferrari",
		KMPerSecond: 140,
	}

	racecars := []Racecar{
		toyota, subaru, porsche, ferrari,
	}

	/* Create the race results channel */

	/* Start the race */

	/* Block until the result are returned */

	timeTaken := raceResults.EndTime.Sub(raceResults.StartTime)
	log.Printf("Race completed in %0.2f seconds!", timeTaken.Seconds())
	log.Printf("Positions:")
	for idx, result := range raceResults.Positions {
		log.Printf("Car %s completed in position %d @%0.2f km/s at %v",
			result.Brand, idx+1, result.KMPerSecond, result.TimeCompleted.UTC().Format(time.RFC3339Nano))
	}
}

func startRace(racecars []Racecar, raceResultsChan chan RaceResults) {
	liveReports := make(map[string][]LiveDriverReport)
	for i := 0; i < len(racecars); i++ {
		liveReports[racecars[i].Brand] = []LiveDriverReport{
			{
				CurrentLap:      0,
				CurrentDistance: 0,
			},
		}
	}

	raceResults := RaceResults{
		StartTime: time.Now(),
	}

	secondsElapsed := 0
	for {
		select {
		/* Add the case to re-calculate where the cars are each second */
			for i := 0; i < len(racecars); i++ {
				racecar := racecars[i]

				log.Printf("\n\nSeconds: %v => Live report: %+v\n\n", secondsElapsed, liveReports)
				lastLiveReport := liveReports[racecar.Brand][secondsElapsed]

				currentDistance := lastLiveReport.CurrentDistance + racecar.KMPerSecond
				currentLap := currentDistance / lapDistanceKM

				liveReports[racecar.Brand] = append(liveReports[racecar.Brand], LiveDriverReport{
					CurrentLap:      currentLap,
					CurrentDistance: currentDistance,
				})

				// Check if this car is done
				if currentLap >= totalLaps && !racecar.Completed {
					raceResults.Positions = append(raceResults.Positions, Result{
						Racecar:       racecar,
						TimeCompleted: time.Now(),
					})

					racecar.Completed = true

					/* Check if all cars have completed racing */
				}
			}

			secondsElapsed++
		}
	}
}

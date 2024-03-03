package helpers

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"network/pkg/models"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func ReadCSV(filePath string) ([]models.Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skipping the header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// create a channel for lines
	lines := make(chan []string)
	// create a channel for errors
	errors := make(chan error)

	// start a goroutine to read lines
	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				close(lines)
				return
			}
			if err != nil {
				errors <- err
				return
			}
			lines <- record
		}
	}()

	// using WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup
	data := make([]models.Data, 0)
	mutex := &sync.Mutex{} // mutex for safe access to shared slice

	// starting worker goroutines
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lines {
				
				country := line[0]

				state := line[1]
				stateId, err := strconv.Atoi(line[2])
				if err != nil {
					errors <- err
					continue
				}

				countryId, err := strconv.Atoi(line[3])
				if err != nil {
					errors <- err
					continue
				}
				networkSpeed, err := strconv.Atoi(line[4])
				if err != nil {
					errors <- err
					continue
				}

				mutex.Lock()
				data = append(data, models.Data{
					Country:      country,
					State:        state,
					StateId: stateId,
					CountryId: countryId,
					NetWorkSpeed: networkSpeed,
				})
				mutex.Unlock()
			}
		}()
	}

	// wait for all workers to finish and close the errors channel
	go func() {
		wg.Wait()
		close(errors)
	}()

	// check for any errors
	for err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func CreateSyntheticDataSet() error {

	var countries = map[string][]string{
		"India":     {"Kerala", "Bangalore", "Delhi", "Hyderabad", "Punjab", "Pune"},
		"China":     {"Hainan", "Gansu", "Shichugan", "Fujian", "Liaonin", "Chethua"},
		"USA":       {"California", "Texas", "Florida", "Alaska", "Hawaii", "Virginia"},
		"Canada":    {"Alberta", "British Columbia", "New Brunswick", "Ontario", "Nova Scotia"},
		"Japan":     {"Hokkaido", "Tohoku", "Kanto", "Chubu", "Kinki", "Chugoku"},
		"Australia": {"New South Wales", "Victoria", "Queensland", "Western Australia", "South Australia", "Tasmania"},
	}

	// Generate fixed IDs for countries and states
	var countryIDs = make(map[string]int)
	var stateIDs = make(map[string]map[string]int)

	for country := range countries {
		countryIDs[country] = rand.Intn(100) + 100
		stateIDs[country] = make(map[string]int)
		for _, state := range countries[country] {
			// Generate fixed IDs for states within each country
			stateIDs[country][state] = rand.Intn(1000) + 1000
		}
	}
	file, err := os.Create("data.csv")
	if err != nil {

		log.Fatal("Cannot create file", err)
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Country", "State", "State ID", "Country ID", "Network Speed"}
	if err := writer.Write(header); err != nil {
		log.Fatal("Error writing header to CSV:", err)
		return err
	}

	// Generate data
	for i := 0; i < 1000; i++ {
		for country, states := range countries {
			countryID := countryIDs[country]
			for _, state := range states {
				stateID := stateIDs[country][state]
				networkSpeed := rand.Intn(1001) // Generates a random number between 0 and 1000
				row := []string{country, state, strconv.Itoa(stateID), strconv.Itoa(countryID), strconv.Itoa(networkSpeed)}
				if err := writer.Write(row); err != nil {
					log.Fatal("Error writing row to CSV:", err)
				}
			}
		}
	}

	log.Println("CSV file generated successfully!")
	return nil
}

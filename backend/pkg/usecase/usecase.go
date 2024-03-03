package usecase

import (
	"fmt"
	helpers "network/pkg/helper"

	"network/pkg/response"
)

func FindAverageSpeedFromData() ([]response.ResponseData, error) {
	data, err := helpers.ReadCSV("./data.csv")
	if err != nil {
		return nil, err
	}

	// Map to store total speed and count for each country
	countrySpeeds := make(map[string]struct {
		TotalSpeed int
		Count      int
		CountryId  int
	})

	// Iterate through the data slice
	for _, entry := range data {
		// Check if the country is already in the map
		if _, exists := countrySpeeds[entry.Country]; !exists {
			// Initialize the map entry with the CountryId
			countrySpeeds[entry.Country] = struct {
				TotalSpeed int
				Count      int
				CountryId  int
			}{
				TotalSpeed: 0,
				Count:      0,
				CountryId:  entry.CountryId, // Set the CountryId here
			}
		}
		current := countrySpeeds[entry.Country]
		current.TotalSpeed += entry.NetWorkSpeed
		current.Count++
		countrySpeeds[entry.Country] = current
	}

	// Slice to store the responses
	var responses []response.ResponseData

	// Calculate the average speed for each country
	for country, speedInfo := range countrySpeeds {
		averageSpeed := float64(speedInfo.TotalSpeed) / float64(speedInfo.Count)
		responses = append(responses, response.ResponseData{
			Country:      country,
			OverallSpeed: int(averageSpeed),
			CountryId:    speedInfo.CountryId,
		})
	}

	return responses, nil
}

func NetworkSpeedInsideCountry(countryId int) ([]response.StateResponse, error) {

	data, err := helpers.ReadCSV("./data.csv")
	if err != nil {
		return nil, err
	}
	stateSpeeds := make(map[string]struct {
		TotalSpeed int
		Count      int
		StateId    int
	})

	for _, entry := range data {
		fmt.Println("HEEYEYYY", countryId, entry.CountryId)

		if countryId == entry.CountryId {
			fmt.Println("HEEYEYYY")

			// Check if the country is already in the map
			if _, exists := stateSpeeds[entry.State]; !exists {
				// Initialize the map entry
				stateSpeeds[entry.State] = struct {
					TotalSpeed int
					Count      int
					StateId    int
				}{}
			}
			current := stateSpeeds[entry.State]
			current.TotalSpeed += entry.NetWorkSpeed
			current.Count++
			current.StateId = entry.StateId
			stateSpeeds[entry.State] = current

		}
	}
	// Slice to store the responses
	var responses []response.StateResponse

	// Calculate the average speed for each country
	for state, speedInfo := range stateSpeeds {
		averageSpeed := float64(speedInfo.TotalSpeed) / float64(speedInfo.Count)
		responses = append(responses, response.StateResponse{
			State:        state,
			OverallSpeed: int(averageSpeed),
			StateId:      speedInfo.StateId,
		})
	}
	return responses, nil
	
}

package handlers

import (
	"encoding/json"
	"fmt"
	"network/pkg/config"
	helpers "network/pkg/helper"
	"network/pkg/response"
	"network/pkg/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAverageSpeedFromData(c *gin.Context) {

	var rd = config.RedisConn
	val, err := rd.Get(c, "country").Result()
	if err == nil {
		// Unmarshal the JSON string into an instance of CountryData
		var countryData []response.ResponseData
		err = json.Unmarshal([]byte(val), &countryData)
		if err != nil {
			response.ErrorResponse(c, 500, "failed to unmarshal data", err, nil)
			return
		}
		fmt.Println("VAAAL", val, "VVV")
		// Now you can work with countryData as a Go struct
		response.SuccessResponse(c, 200, "successfully fetched using Redis", countryData)
		return
	}

	data, err := usecase.FindAverageSpeedFromData()
	if err != nil {
		response.ErrorResponse(c, 500, "oops something wrong", err, nil)
		return
	}

	dataString, err := json.Marshal(data)
	if err != nil {
		response.ErrorResponse(c, 500, "failed to serialize data", err, nil)
		return
	}

	// Store the serialized data in Redis
	err = rd.Set(c, "country", dataString, 0).Err()
	if err != nil {
		response.ErrorResponse(c, 500, "failed to store data in Redis", err, nil)
		return
	}
	response.SuccessResponse(c, 200, "succesfully fetched the average speed", data)

}

func CreateDataSet(c *gin.Context) {
	if err := helpers.CreateSyntheticDataSet(); err != nil {
		response.ErrorResponse(c, 500, "oops something wrong", err, nil)
		return
	}
	response.SuccessResponse(c, 200, "succesfully created data set", nil)
}

func NetworkSpeedInStates(c *gin.Context) {
	id := c.Param("id")
	// type Country struct {
	// 	Country string `json:"country"`
	// }
	// var countryIp Country

	// if err := c.ShouldBindJSON(&countryIp); err != nil {
	// 	response.ErrorResponse(c, 400, "error in binding json", err, nil)
	// 	return
	// }
	// Id, _ := strconv.Atoi(countryIp.Country)

	var rd = config.RedisConn
	val, err := rd.Get(c, id).Result()
	if err == nil {
		// Unmarshal the JSON string into an instance of CountryData
		var stateData []response.StateResponse
		err = json.Unmarshal([]byte(val), &stateData)
		if err != nil {
			response.ErrorResponse(c, 500, "failed to unmarshal data", err, nil)
			return
		}
		// Now you can work with countryData as a Go struct
		response.SuccessResponse(c, 200, "successfully fetched using Redis", stateData)
		return
	}
	Id, _ := strconv.Atoi(id)
	data, err := usecase.NetworkSpeedInsideCountry(Id)
	if err != nil {
		response.ErrorResponse(c, 400, "oops someting went wrong", err, nil)
		return
	}

	dataString, err := json.Marshal(data)
	if err != nil {
		response.ErrorResponse(c, 500, "failed to serialize data", err, nil)
		return
	}

	// Store the serialized data in Redis
	err = rd.Set(c, id, dataString, 0).Err()
	if err != nil {
		response.ErrorResponse(c, 500, "failed to store data in Redis", err, nil)
		return
	}
	response.SuccessResponse(c, 200, "succesfully created data set", data)

}

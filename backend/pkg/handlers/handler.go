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

	rd, err := config.GetRedisConnection()

	if err != nil {
		response.ErrorResponse(c, 500, "oops something wrong", err, nil)
		return
	}

	val, err := rd.Get(c, "country").Result()
	fmt.Println("VAL", val)
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
	type Country struct {
		Country string `json:"country"`
	}
	var a Country

	if err := c.ShouldBindJSON(&a); err != nil {
		response.ErrorResponse(c, 400, "error in binding json", err, nil)
		return
	}
	Id, _ := strconv.Atoi(a.Country)
	fmt.Println(a, "ID", Id)
	data, err := usecase.NetworkSpeedInsideCountry(Id)
	if err != nil {
		response.ErrorResponse(c, 400, "oops someting went wrong", err, nil)
		return
	}
	response.SuccessResponse(c, 200, "succesfully created data set", data)

}

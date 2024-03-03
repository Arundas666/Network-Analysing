package handlers

import (
	"fmt"
	helpers "network/pkg/helper"
	"network/pkg/response"
	"network/pkg/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAverageSpeedFromData(c *gin.Context) {

	data, err := usecase.FindAverageSpeedFromData()
	if err != nil {
		response.ErrorResponse(c, 500, "oops something wrong", err, nil)
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
	fmt.Println(a,"ID", Id)
	data, err := usecase.NetworkSpeedInsideCountry(Id)
	if err != nil {
		response.ErrorResponse(c, 400, "oops someting went wrong", err, nil)
		return
	}
	response.SuccessResponse(c, 200, "succesfully created data set", data)

}

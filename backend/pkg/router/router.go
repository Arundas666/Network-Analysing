package router

import (
	"network/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.GET("/find-avg", handlers.FindAverageSpeedFromData)
	r.GET("/create-data", handlers.CreateDataSet)

	r.POST("/avg-in-states", handlers.NetworkSpeedInStates)

}

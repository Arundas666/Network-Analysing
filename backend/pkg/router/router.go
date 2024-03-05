package router

import (
	"network/pkg/handlers"
	"network/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.Group("/", middleware.CorsMiddleware)
	{

		r.GET("/find-avg", handlers.FindAverageSpeedFromData)
		r.GET("/avg-in-states", handlers.NetworkSpeedInStates)
	}
	r.GET("/create-data", handlers.CreateDataSet)

}

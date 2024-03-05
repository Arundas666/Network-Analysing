package router

import (
	"network/pkg/handlers"
	"network/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.GET("/find-avg", middleware.CorsMiddleware, handlers.FindAverageSpeedFromData)
	r.GET("/avg-in-states", middleware.CorsMiddleware, handlers.NetworkSpeedInStates)

	r.GET("/create-data", handlers.CreateDataSet)

}

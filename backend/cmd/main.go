package main

import (
	"fmt"
	"log"
	"network/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	router.Router(route.Group("/"))

	if err := route.Run(":8080"); err != nil {
		fmt.Println("ERR", err)
		log.Fatalf("Error starting server on %s: %v", err)
	}
}

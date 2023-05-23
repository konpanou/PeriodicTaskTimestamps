package main

import (
	"fmt"
	"log"
	"os"

	"PeriodicTaskTimestamps/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read host address and port from command line arguments
	host := os.Args[1]
	port := os.Args[2]

	// Create a new Gin router
	router := gin.Default()

	// Define the /ptlist endpoint
	router.GET("/ptlist", api.GetTimestamps)

	// Start the server
	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Starting server at %s", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	router "janburzinski.de/unilearnapp/routes"
)

func main() {
	r := router.SetupRoutes()

	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if parsedPort, err := strconv.Atoi(portStr); err == nil {
			port = parsedPort
		}
	}

	log.Println("Running on Port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

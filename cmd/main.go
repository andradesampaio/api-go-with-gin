package main

import (
	"api-go-with-gin/internal/router"
	"api-go-with-gin/internal/database"
	)

func main() {
	database.Connect()
	router.HandleRequests()
}


package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.ConnectDB() 
	data.ConnectUserDB() 

	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

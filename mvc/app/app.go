package app

import (
	"log"
	"mvc/controllers"
	"net/http"
)

func StartApp(){

	// Users
	http.HandleFunc("/users", controllers.GetUser)

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

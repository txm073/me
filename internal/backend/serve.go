package me

import (
	"log"
	"net/http"
)

func RunServer() {
	log.Println("Server running on port", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

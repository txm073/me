package me

import (
	"log"
	"net/http"
)

func RunServer() {
	log.Println("Server running on port", Address)
	log.Fatal(http.ListenAndServe(Address, nil))
}

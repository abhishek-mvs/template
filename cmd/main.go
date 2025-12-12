package main

import (
	"template/api/rest"
	"log"
)

func main() {
	server := rest.HttpBuildServer()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on port 8080")
}

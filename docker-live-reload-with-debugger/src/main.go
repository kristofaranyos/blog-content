package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world!"))
	})

	log.Println("starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(fmt.Errorf("could not start server: %w", err))
	}
}

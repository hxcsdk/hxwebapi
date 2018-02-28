package main

import (
	"log"
	"net/http"
)

const (
	// the listening port
	port = ":80"
	//port = ":9155"
)

func main() {
	service := NewService()
	log.Println("hxwebapi started on", port)
	log.Fatal(http.ListenAndServe(port, service.Router))
}

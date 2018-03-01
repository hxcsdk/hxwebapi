package main

import (
	"log"
	"net"
	"net/http"
)

const (
	// the listening port
	port = ":8088"
	//port = ":9155"
)

func main() {
	service := NewService()
	log.Println("hxwebapi started on", port)
	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, service.Router))
}

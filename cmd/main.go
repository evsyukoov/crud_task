package main

import (
	"AvitoTest/app"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", app.Handler);
	if http.ListenAndServe(":80", nil) != nil {
		log.Fatal("Can't bind on this adress and port, stop!");
	}
}

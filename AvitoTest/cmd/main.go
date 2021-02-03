package main

import (
	"AvitoTest/app"
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Printf("\033[0;32mServer is running at 127.0.0.1:80\n\033[0m");
	http.HandleFunc("/", app.Handler);
	if http.ListenAndServe(":80", nil) != nil {
		log.Fatal("Can't bind on this adress and port, stop!");
	}
}

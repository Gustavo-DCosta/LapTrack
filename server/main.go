package main

import (
	"github.com/Gustavo-DCosta/LapTrack/router"
	"net/http"
	"log"
)

func main() {
	router.Routing()
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/cookie_sync/cmd/lib/server"
)

var (
	p              = os.Getenv("PORT_NUMBER_OF_SERVER")
	handlerFuncMap = map[string]func(w http.ResponseWriter, r *http.Request){
		// "delivery":,
		// "click":,
		// "cv":
	}
)

func main() {
	s := server.New(handlerFuncMap)
	if err := s.Start(p); err != nil {
		log.Fatal(err)
	}

}

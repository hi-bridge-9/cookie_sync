package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	path map[string]func(http.ResponseWriter, *http.Request)
}

func New(path map[string]func(http.ResponseWriter, *http.Request)) *Server {
	return &Server{
		path: path,
	}
}

func (s *Server) Start(port string) error {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	log.Printf("[%v]server is started\n", time.Now().Local())

	if port == "" {
		// set default port number
		port = "80"
	}
	for path, handler := range s.path {
		http.HandleFunc(path, handler)
	}

	return http.ListenAndServe(":"+port, nil)
}

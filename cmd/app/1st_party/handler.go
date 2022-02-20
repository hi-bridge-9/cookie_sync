package main

import (
	"fmt"
	"net/http"
)


func mediaPageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("Invalid request method: %v", r.Method)
	}

	return nil
}

func deliveryHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("Invalid request method: %v", r.Method)
	}

	return nil
}

func clickHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("Invalid request method: %v", r.Method)
	}

	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := TimeResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}
	writer.Write(jsonResponse)
}

func main() {
	port := `:8795`

	http.HandleFunc("/time", Get)
	fmt.Println("Launching an http server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Unexpected error occurred while launching an http server:", err)
	}
}

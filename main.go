package main

import (
	"net/http"
	"fmt"
	"Environment/server"
)

func main() {
	fmt.Println("Server active at :8080")
	http.HandleFunc("/", server.HomeView)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("No go : 500")
		return
	}
}
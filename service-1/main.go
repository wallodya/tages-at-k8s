package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Some random string")
}

func main() {
	http.HandleFunc("/service-1", mainHandler)
	fmt.Println("Server listens on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

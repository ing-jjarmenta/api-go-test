package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	log.Println("Server 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("Error desde el serve")
		log.Fatal(err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hola desde la API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

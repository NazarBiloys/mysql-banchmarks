package main

import (
	"github.com/NazarBiloys/mysql-banchmarks/internal/service"
	"log"
    "fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := service.MakeUser(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":90", nil))
}

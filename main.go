package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Totiruzi/dogma-crud-docker/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on port 8080....")

	log.Fatal(http.ListenAndServe(":8080", r))
}
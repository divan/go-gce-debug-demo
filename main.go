package main

import (
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	str := "Hello, World"
	str = strings.Replace(str, "Hello", "Hola", -1)
	str = strings.Replace(str, "World", "Mundo", -1)
	response := []byte(str)
	w.Write(response)
}

func main() {
	http.HandleFunc("/", hello)

	log.Println("Starting HTTP server on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

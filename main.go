package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, Kubernetes com Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor ouvindo na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

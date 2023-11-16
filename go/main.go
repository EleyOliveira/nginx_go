//package = compila como um executável
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Cria a primeira rota
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá Osasco")
	})

	//Cria uma rota que recebe um parâmetro URL
	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greet/"):]
		fmt.Fprintf(w, "Olá %s\n", name) 
	})

	//Instancia o servidor
	http.ListenAndServe(":9990", nil)
}
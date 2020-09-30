package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	age    int
	name   string
	prenom string
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html>")
	fmt.Fprint(w, "<head>")
	fmt.Fprint(w, "</head>")
	fmt.Fprint(w, "<body>")
	fmt.Fprint(w, "<h1>La météo est de 25 degrés<H1>")
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")
	p := Person{age: 12, prenom: "Glenn", name: "Louarn"}
	WriteJSON(p)

}

func main() {
	http.HandleFunc("/temp", tempHandler)
	err := http.ListenAndServe(":8888", nil)

	log.Fatal(err)
}

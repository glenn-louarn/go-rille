package main

import (
	"github.com/rs/cors"
	"net/http"
)

type person struct {
	Age    int
	Name   string
	Prenom string
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
	/*p := person{Age: 12, Prenom: "Glenn", Name: "Louarn"}
	json.WriteJSON(w, p)*/
}

func main() {
	/*http.HandleFunc("/temp", tempHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8081", handler)
}

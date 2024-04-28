package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /habitos", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("Obtiene lista de habitos")))
	})
	router.HandleFunc("GET /habito/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(("id"))

		fmt.Fprintf(w, "nota by id: %s", id)
	})

	server := &http.Server{
		Addr:    ":6060",
		Handler: router,
	}

	log.Println("Servidor esperando en el puerto ", server.Addr)

	server.ListenAndServe()

}

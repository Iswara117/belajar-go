package main

import (
	"sesi10/menu"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
    // proses pembuatan router oleh chi-router
	router := chi.NewRouter()

    // untuk database, kita isi sebagai nil terlebih dahulu
	menu.Register(router, nil)

	const port = ":8000"

	log.Println("server running at port", port)

    // proses listen port dengan custom router
	http.ListenAndServe(port, router)
}
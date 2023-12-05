package main

import (
	// "sesi10/menu"
	"log"
	"net/http"
	"os"
	"sesi11/database"
	"sesi11/employee"
	"sesi11/menu"
	"sesi11/transactions"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


func main() {
    // proses mencoba connect ke database
    // ========== UPDATE ==========
	db, err := database.ConnectPostgres(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
    // cek apakah ada error atau engga
	if err != nil {
		panic(err)
	}

    // cek apakah databasenya nil atau engga
    // kalau nil, berarti gagal membuat object database nya
    if db == nil {
        panic("db not connected")
    }
    // ================
	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Atur sumber-sumber yang diizinkan (bisa diganti dengan domain yang spesifik)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Atur durasi caching CORS Preflight requests (opsional)
	}))
//     // untuk database, kita isi sebagai nil terlebih dahulu
	menu.Register(router, db)
	employee.Register(router, db)
	transactions.Register(router, db)

	const port = ":8000"

	log.Println("server running at port", port)

    // proses listen port dengan custom router
	http.ListenAndServe(port, router)
    // code router sebelumnya
}	

// main old
// func main() {
//     // proses pembuatan router oleh chi-router
// 	router := chi.NewRouter()

//     // untuk database, kita isi sebagai nil terlebih dahulu
// 	menu.Register(router, nil)

// 	const port = ":8000"

// 	log.Println("server running at port", port)

//     // proses listen port dengan custom router
// 	http.ListenAndServe(port, router)
// }
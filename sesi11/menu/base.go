package menu

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// function ini untuk melakukan init terhadap semua
// hal yang dibutuhkan oleh Menu Services
func Register(router *chi.Mux, db *sql.DB) {
	
	repo := NewRepository(db) //{}
    // inject into services layer
	// fmt.Print(repo)
	svc := NewService(repo) // {{}}
	// fmt.Print(svc)

	// println(svc)
	handler := NewHandler(svc) //{{{}}}
	// fmt.Print(handler)

	// println(handler)
    // seperti grouping endpoint
    // jadi yang ada di dalamnya sudah memiliki endpoint /api/menus 
    // sebagai dasarnya
	
	router.Route("/api/menus", func(r chi.Router) {

        // endpoint : /api/menus
		r.Get("/", handler.GetAll)
        
        // endpoint : /api/menus/{id}
		r.Get("/{id}", handler.GetById)

        // endpoint : /api/menus
		r.Post("/", handler.Create)
		r.Put("/{id}", handler.UpdateById)
	})
}
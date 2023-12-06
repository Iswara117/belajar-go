package auth

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"sesi11/infra/router/chi"
)

// function ini untuk melakukan init terhadap semua
// hal yang dibutuhkan oleh Auth Services
func Register(router *chi.Mux, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

    // seperti grouping endpoint
    // jadi yang ada di dalamnya sudah memiliki endpoint /api/auth 
    // sebagai dasarnya
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/signup", handler.Register)
		r.Post("/signin", handler.Login)
		r.Group(func(r chi.Router) {
            // use middleware selalu di awal 
			r.Use(routerChi.CheckToken)
			r.Get("/profile", handler.Profile)
		})
	})
}
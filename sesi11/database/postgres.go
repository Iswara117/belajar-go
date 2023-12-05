package database

import (
	"database/sql"
	"fmt"

	// ini harus ditambahkan
	// untuk melakukan registrasi terhadap driver postgresnya
	_ "github.com/lib/pq"
)

func ConnectPostgres(host, port, user, pass, dbname string) (db *sql.DB, err error) {
	// membuat data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "admin", "restaurant",
	)

	// melakukan open koneksi ke postgres
	// driver postgres bisa di dapat dari melakukan import
	// dengan cara import _ "github.com/lib/pq"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	// validate if db berhasil untuk terhubung
	// dengan cara melakukan `ping`
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

package config

import (
    // driver SQL bawaan dari go
	"database/sql"
	"fmt"
	"time"

    // package yang sudah kita install sebelumnya.
    // ini berfungsi untuk menambahkan postgres dalam package dabase/sql
	_ "github.com/lib/pq"
)

// berfungsi untuk membuat koneksi
func CreateConnection() *sql.DB {

    // silahkan sesuaikan dengan yang kamu punya
    // normalnya untuk host = localhost
	host := "localhost"

    // port biasanya 5432, tergantung dengan yang kamu buat saat proses installasi
	port := "5432"

    // username dari database kamu
	user := "postgres"

    // password dari database kamu
	pass := "admin"

    // nama database
	dbname := "restaurant"

    // proses pembuatan koneksi
	dbs, err := getPostgres(host, port, user, pass, dbname)
	if err != nil {
		panic(err)
	}

	err = dbs.Ping()

	if err != nil {
		panic(err)
	}

	return dbs

}

func getPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	// membuat data source
	desc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbs, err := createConnection(desc)
	if err != nil {

		return nil, err
	}

	return dbs, nil
}

func createConnection(desc string) (*sql.DB, error) {
	// proses membuka koneksi
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}

	// method ini berfungsi untuk melakukan set minimum koneksi yang dibuat
	// jadi saat program dijaklankan, dia akan membuat 10 koneksi yang pada posisi idle
	db.SetMaxIdleConns(10)

	// method ini berfungsi untuk melakukan set maximum jumlah koneksi yang dibuat
	db.SetMaxOpenConns(25)

	// method ini berfungsi untuk menentukan masa aktif sebuah koneksi saat posisi idle
	// jika melebih batas waktu yang ditentukan, maka koneksi akan dihapus hingga batas SetMaxIdleConns
	db.SetConnMaxIdleTime(5 * time.Minute)

	// method ini berfungsi untuk menentukan lamanya sebuah koneksi itu ada
	// jika lewat dari waktunya, maka koneksi akan dihapus dan di generate ulang
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
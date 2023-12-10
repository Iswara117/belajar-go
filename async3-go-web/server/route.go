package server

import (
	"database/sql"
	"fmt"
	"async3-go-web/server/controllers"
	"net/http"
)

// fungsi ini akan di panggil di main.go
// jadi kita akan membuatnya sebagai Exported.
// fungsi ini juga berguna untuk wrap semua rute dan sekaligus
// menjalankan servernya.
// pada fungsi ini, kita membutuhkan sebuah router, port, dan db yang 
// akan diinject pada main.go nantinya
func StartServer(router *http.ServeMux, port string, db *sql.DB) {
    // proses pemanggilan route
	buildRoute(router, db)


	fmt.Println("Server running at", port)
    // proses running server
	http.ListenAndServe(port, router)
}

// fungsi ini berguna untuk wrapping rute.
// fungsi ini dibuat Unexported, karena akan dipanggil pada StartServer
func buildRoute(router *http.ServeMux, db *sql.DB) {
	homeRoute(router)
	employeeRoute(router, db)
}

func employeeRoute(router *http.ServeMux, db *sql.DB) {

    // proses inisialisasi controller, dan menginject db kedalamnya
	employeeController := controllers.NewEmployeeController(db)

    // beberapa endpoint yang kita gunakan, sesuai dengan API Spec
	router.HandleFunc("/employees", employeeController.Index)
	router.HandleFunc("/employees/update", employeeController.UpdateByID)
	router.HandleFunc("/employees/add", employeeController.Add)
	router.HandleFunc("/employees/delete", employeeController.DeleteByID)
}

// fungsi untuk handle rute rute dari home controller
func homeRoute(router *http.ServeMux) {
	// proses inisialisasi pembuatan home controller.
	// ini berfungsi agar bisa menggunakan methodnya
	homeController := controllers.NewHomeController()

	// ini berfungsi untuk routing nya.
	// "/" artinya => http://localhost:port/
	// params ke 2 berguna untuk memanggil method apa yang akan dijalankan, jika
	// path tadi di panggil.
	router.HandleFunc("/", homeController.Index)

}
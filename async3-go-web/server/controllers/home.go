package controllers

import (
	"async3-go-web/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

// membuat interface, untuk menampung
// semua method yang akan digunakan.
type HomeController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

// fungsi ini digunakan agar bisa melakukan implementasi
// terhadap method pada interface tadi.
func NewHomeController() HomeController {
	return &homeController{}
}

// method Index
func (*homeController) Index(w http.ResponseWriter, r *http.Request) {
    // proses mengambil template pada folder `static`.
	// template.ParseFiles() berfungsi untuk mengambil filenya
	// fungsi ini membutuhkan parameter berupa ...path dari file file html-nya.
	tmpl, err := template.ParseFiles(path.Join("static", "pages/home/index.html"), utils.LayoutMaster)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// proses render file yang telah kita panggil diatas.
	// method Execute membutuhkan 2 parameter, yaitu sebuah ResponseWriter dan sebuah data.
	// karnea pada method Index ini kita tidak membutuhkan data, maka cukup ditulis dengan nil
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
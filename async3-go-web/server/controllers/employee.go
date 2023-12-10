// server/controllers/employee.go

package controllers

import (
	apps "async3-go-web/server/apps/web"
	params "async3-go-web/server/params/employee"
	"async3-go-web/server/services"
	"async3-go-web/server/utils"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	// "github.com/google/uuid"
)

type EmployeeController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type employeeController struct {
	DB *sql.DB
}

func NewEmployeeController(db *sql.DB) EmployeeController {
	return &employeeController{
		DB: db,
	}
}

// berfungsi untuk menghandle halaman awal dari employee
func (e *employeeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // ====== TAMBAHAN MULAI DARI SINI ======

    // proses menggunakan method GetAllEmployees dari services layer
    // method ini akan mengembalikan nilai slice
	employees := services.NewEmployeeService(e.DB).GetAllEmployees()

    // lalu memasukkannya kedalam Data, agar bisa di render di template HTML nya
	web := apps.RenderWeb{
        Title: "Halaman Employee",
		Data:  employees,
	}
    // ====== END ======   

	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
// berfungsi untuk menampilkan form dan melakukan proses tambah data ke database
func (e *employeeController) Add(w http.ResponseWriter, r *http.Request) {
    // mengambil method apa yang sedang aktif pada endpoint ini
	method := r.Method  

    // jika methodnya GET, maka akan menampilkan form tambah employee
	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/add.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

        // karena hanya menampilkan form biasa, maka tidak perlu menambahkan data didalamnya
		web := apps.RenderWeb{
			Title: "Tambah Pegawai",
		}

		err = tmpl.Execute(w, web)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if method == "POST" {
        // dan jika methodnya adalah POST, maka lakukan proses insert ke database

        // ini berfungsi untuk parsing form input dari html
		fmt.Println("Aku Masuk Request Post")
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

        // proses pengambilan nilai input dari html.
        // nip, address, dan name itu akan didapat dari tag input di properti `name` pada html
        // contohnya : <input name="nip" />
        // atau <input name="address" />
		var request params.EmployeeCreate = params.EmployeeCreate{
			NIP:     r.Form.Get("nip"),
			Address: r.Form.Get("address"),
			Name:    r.Form.Get("name"),
		}

        // setelah berhasil membuat sebuah request yang bertipe data params, 
        // maka selanjutnya adalah proses mengirim data tersebut ke services layer.
        // Inisialisasi services layer
		employeeServices := services.NewEmployeeService(e.DB)
		fmt.Println(employeeServices)

        // isSuccess akan bernilai true atau false. Jadi kita akan melakukan pengecekan
		isSuccess := employeeServices.CreateNewEmployee(&request)

		msg := ""

        // jika success, maka tampilkan alert dan redirect menggunakan javascript
		if isSuccess {
			msg = `
				<script>
					alert("Tambah data pegawai berhasil !")
					window.location.href="../employee"
				</script>
			`
		} else {

			msg = `
				<script>
					alert("Tambah data pegawai gagal !")
					window.location.href="../employee"
				</script>
			`
		}

        // karena ini hanyalah sebuah proses untuk menambahkan data,
        // maka kita tidak perlu merender sebuah tampilan web.
        // cukup mengembalikan msg tadi, maka redirect nya akan di handle
        // sama msg.
		w.Write([]byte(msg))
	} else {
        // selain itu, method apapun dilarang
		msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
		w.Write([]byte(msg))
	}
}

// berfungsi untuk menghapus data employee berdasarkan ID nya
func (e *employeeController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	query := r.URL.Query()

    // mengambil id pada query
	id := query.Get("id")

	intID, _ := strconv.Atoi(id)
	if method == "GET" {

		tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/update.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get data awal dari si pegawai, berdasarkan IDnya. Jadi data ini akan di update nantinya
		employee := services.NewEmployeeService(e.DB).GetEmployeeByID(intID)

		web := apps.RenderWeb{
			Title: "Halaman Detail Employee",
			Data:  employee,
		}

		err = tmpl.Execute(w, web)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// id yang di dapat akan di parse kedalam bentuk UUID.
		// karena pada params, ia membutuhkan ID berupa UUID.
		// newID, err := uuid.Parse(id)
		// if err != nil {
		// 	log.Println(err)
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		var request params.EmployeeUpdate = params.EmployeeUpdate{
			ID:      intID,
			NIP:     r.Form.Get("nip"),
			Address: r.Form.Get("address"),
			Name:    r.Form.Get("name"),
		}

		employeeServices := services.NewEmployeeService(e.DB)

		isSuccess := employeeServices.UpdateByID(&request)

		msg := ""
		if isSuccess {
			msg = `
				<script>
					alert("Ubah data pegawai berhasil !")
					window.location.href="../employees"
				</script>
			`
		} else {

			msg = `
				<script>
					alert("Ubah data pegawai gagal !")
					window.location.href="../employees"
				</script>
			`
		}

		w.Write([]byte(msg))
	} else {
		msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
		w.Write([]byte(msg))
	}
}

// berfungsi untuk mengubah data employee berdasarkan ID nya
func (e *employeeController) DeleteByID(w http.ResponseWriter, r *http.Request) {
    // proses penangkapan query
	query := r.URL.Query()

    // mengambil id pada query
    // ini bisa juga dilakukan dengan query.Get("id")
	id := query["id"][0]
	intID, _ := strconv.Atoi(id)
    // proses deleteData
	deleteData := services.NewEmployeeService(e.DB).DeleteEmbloyeeByID(intID)

	msg := ""
	if deleteData {
		msg = `
			<script>
				alert("Hapus data pegawai berhasil !")
				window.location.href="../employees"
			</script>
		`
	} else {

		msg = `
			<script>
				alert("Hapus data pegawai gagal !")
				window.location.href="../employees"
			</script>
		`
	}

	w.Write([]byte(msg))
}
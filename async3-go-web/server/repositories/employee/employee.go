package repositories

import "async3-go-web/server/models"

type EmployeeRepository interface {
    // ini berfungsi untuk menyimpan data ke database
	Save(employee *models.Employee) error

    // berfungsi untuk mengambil semua data employee
	FindAll() (*[]models.Employee, error)

    // berfungsi untuk mengambil data employee berdasarkan ID
	FindByID(id int) (*models.Employee, error)

    // berfungsi untuk mengubah data employee
	UpdateByID(employee *models.Employee) error

    // berfungsi untuk menghapus data pegawai berdasarkan ID
	DeleteByID(id int) error
}
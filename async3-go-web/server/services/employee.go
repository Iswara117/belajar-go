package services

import (
	"async3-go-web/server/helper"
	"async3-go-web/server/models"
	params "async3-go-web/server/params/employee"
	repositories "async3-go-web/server/repositories/employee"
	"database/sql"
	"fmt"
	"time"
)

// disini kita akan menambahkan repositories tadi
type EmployeeServices struct {
	EmployeeRepository repositories.EmployeeRepository
	DB                 *sql.DB
}

// inisialisasi employee services
func NewEmployeeService(db *sql.DB) *EmployeeServices {
    // proses inisialisasi repositories dengan menginject DB
	repositories := repositories.NewEmployeeRepository(db)
	return &EmployeeServices{
		DB:                 db,
		EmployeeRepository: repositories,
	}
}


func (e *EmployeeServices) CreateNewEmployee(request *params.EmployeeCreate) bool {
	fmt.Println("masuk guys")
    // proses recover harus dilakukan dengan defer
	defer helper.HandleError()

    // melakukan parse data dari yang berupa params menjadi model
	emp := request.ParseToModel()
	fmt.Println(emp)

    // memanggil method Save pada repositories, dan mengirimkan 
    // data yang sudah berbentuk models
	err := e.EmployeeRepository.Save(emp)

	if err != nil {
		helper.HandlePanicIfError(err)
		return false
	}

	return true
}

func (e *EmployeeServices) GetAllEmployees() *[]helper.EmployeeSingleView {
	defer helper.HandleError()

    // proses pemanggilan method FindAll()
    // method ini akan mereturn slice of models dan error
	employees, err := e.EmployeeRepository.FindAll()

	helper.HandlePanicIfError(err)

	return makeEmployeeListView(employees)

}

// fungsi tambahan, untuk melakukan proses parsing data dari models ke params
// fungsi ini akan mereturn data dalam bentuk slice
func makeEmployeeListView(models *[]models.Employee) *[]helper.EmployeeSingleView {
	var params []helper.EmployeeSingleView
	for _, model := range *models {
        // menggunakan method makeEmployeeSingleView agar dapat
        // sebagai tempat parsing datanya
		params = append(params, *makeEmployeeSingleView(&model))
	}
	return &params
}

// fungsi tambahan, untuk melakukan proses parsing dari data models ke params
// fungsi ini akan mereturn singleView
func makeEmployeeSingleView(models *models.Employee) *helper.EmployeeSingleView {
	return &helper.EmployeeSingleView{
		ID:        models.ID,
		NIP:       models.NIP,
		Name:      models.Name,
		Address:   models.Address,
		CreatedAt: models.CreatedAt.Format(time.RFC3339),
		UpdatedAt: models.UpdatedAt.Format(time.RFC3339),
	}
}

func (e *EmployeeServices) DeleteEmbloyeeByID(id int) bool {
	defer helper.HandleError()
	err := e.EmployeeRepository.DeleteByID(id)
	if err != nil {
		helper.HandlePanicIfError(err)
		return false
	}

	return true
}

func (e *EmployeeServices) GetEmployeeByID(id int) *helper.EmployeeSingleView {
	defer helper.HandleError()

    // proses pemanggilan method FindByID
    // method ini akan mengembalikan dalam bentuk models
	// intID, err := strconv.Atoi(id)
	employee, err := e.EmployeeRepository.FindByID(id)
	helper.HandlePanicIfError(err)

    // data dalam bentuk models tadi akan di parsing menjadi bentuk params
    // agar bisa di consume sama controllers
	return makeEmployeeSingleView(employee)
}


func (e *EmployeeServices) UpdateByID(request *params.EmployeeUpdate) bool {
	defer helper.HandleError()

    // data params dari controllers akan di parse dulu agar menjadi data models
	model := request.ParseToModel()

    // memanggil methid UpdateByID.
    // method ini akan mengembalikan sebuah error
	err := e.EmployeeRepository.UpdateByID(model)
	helper.HandlePanicIfError(err)

    // jika success, maka akan mereturn true
	return true

}
package params

import (
	"async3-go-web/server/models"
	"time"
	// "github.com/google/uuid"
)

// berfungsi untuk handle request
// untuk membuat employee baru
type EmployeeCreate struct {
	NIP     string
	Name    string
	Address string
}

// method ini berfungsi untuk mengubah data params ke models
func (e *EmployeeCreate) ParseToModel() *models.Employee {
	employee := models.NewEmployee()
	employee.Address = e.Address
	employee.Name = e.Name
	employee.NIP = e.NIP
	return employee
}

// ini berfungsi untuk handle request
// jika ada update employee
type EmployeeUpdate struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	UpdatedAt time.Time
}

// method ini berfungsi untuk mengubah data params ke models
func (e *EmployeeUpdate) ParseToModel() *models.Employee {
	return &models.Employee{
		ID:        e.ID,
		NIP:       e.NIP,
		Name:      e.Name,
		Address:   e.Address,
		UpdatedAt: time.Now(),
	}
}
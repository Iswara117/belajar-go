package models

import (
	"time"

	// "github.com/google/uuid"
)

// properti yang ada diambil dari table employees di Database
type Employee struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ini berfungsi saat proses membuat model employee baru, maka
// secara default akan langsung men-generate ID, CreatedAt, dan UpdatedAt
func NewEmployee() *Employee {
	return &Employee{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
package employee

import "time"

type Employee struct {
	Id        int       `json:"id"`
	Nip       string    `json:"nip"`
	Name  	  string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func New(name string, nip string, address string) Employee {
	return Employee{
		Name:      name,
		Nip: 	   nip,
		Address:   address,		  	
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
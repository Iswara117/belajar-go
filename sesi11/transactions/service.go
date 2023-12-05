package transactions

import (
	"database/sql"
	"log"
	"errors"
)

// "fmt"
// "log"

// interface dimiliki oleh consumer
// jadi peletakan interface ada di level siapa yang menggunakannya
type RepositoryInterface interface {
	Save(data Transactions) (err error)
	GetMenuById(id int) (menu Menu, err error)
	GetEmployeeById(id int) (emp Employee, err error)
	// UpdateById(id int, menu Menu ) (err error)
	GetAll() (data []Transactions, err error)
	FindById(id int) (data Transactions, err error)
}

type Service struct {
	// membutuhkan dependency ke repository
	// yang mana harus sesuai dengan kontrak yang sudah di sepakati (interface)
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Pay(req Transactions) (err error) {
	_, err = s.repo.GetEmployeeById(req.EmployeeId)
	if err != nil {
		// log.Println("error when try to GetEmployeeById with emp_id:", req.Employee.Id, "and error:", err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}
	menu, err := s.repo.GetMenuById(req.MenuId)
	if err != nil {
		log.Println("error when try to GetMenuById with menu_id:", req.MenuId, "and error", err.Error())
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}
	req.SetTotal(menu.Price)

	err = s.repo.Save(req)
	return
}
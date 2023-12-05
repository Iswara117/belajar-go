package employee

import (
	"log"
	// "fmt"
)

type RepositoryInterface interface {
	Save(employee Employee) (err error)
	GetAll() (employees []Employee, err error)
	GetById(id int) (employee Employee, err error)
	DeleteById(id int) (err error)
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

func (s Service) Save(employee Employee) (err error) {
	// disini kita memanggil repository nya
	// untuk  melakukan Save data
	// fmt.Print(menu)
	err = s.repo.Save(employee)
	if err != nil {
		log.Println("error when try to save menu with error", err.Error())
		return
	}
	return
}
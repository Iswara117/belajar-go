package menu

import (
	"fmt"
	"log"
)

// interface dimiliki oleh consumer
// jadi peletakan interface ada di level siapa yang menggunakannya
type RepositoryInterface interface {
	Save(menu Menu) (err error)
	GetAll() (menus []Menu, err error)
	GetById(id int) (menu Menu, err error)
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
// log.Print(Service)


func (s Service) Save(menu Menu) (err error) {
    // disini kita memanggil repository nya 
    // untuk  melakukan Save data
	fmt.Print(menu)
	err = s.repo.Save(menu)
	if err != nil {
		log.Println("error when try to save menu with error", err.Error())
		return
	}
	return
}

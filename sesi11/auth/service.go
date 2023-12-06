package auth

import (
	"sesi11/utility"
	"database/sql"
	"fmt"
	"log"
)

// interface dimiliki oleh consumer
// jadi peletakan interface ada di level siapa yang menggunakannya
type RepositoryInterface interface {
	Create(auth Auth) (err error)
	GetByEmail(email string) (auth Auth, err error)
	GetById(id interface{}) (auth Auth, err error)
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

func (s Service) Create(auth Auth) (err error) {
	// hash password user sebelum di insert ke db
	auth.Password, err = utility.Hash(auth.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	// insert ke datasource
	err = s.repo.Create(auth)
	if err != nil {
		log.Println("error when try to create auth with error", err.Error())
		return
	}
	return
}

func (s Service) Login(req Auth) (auth Auth, err error) {
	// check apakah user dengan email tersebut ada atau tidak
	auth, err = s.repo.GetByEmail(req.Email)
	if err != nil {
		log.Println("error when try to get auth by email with error", err.Error())
		if err == sql.ErrNoRows {
			err = fmt.Errorf("username or password not found")
			return
		}
		return
	}

	// lakukan verifikasi
	// password yang dari database, dalam bentuk hash
	// jadi perlu kita verify
	err = utility.Verify(auth.Password, req.Password)
	if err != nil {
		log.Println("error when try to verify password with error", err.Error())
		err = fmt.Errorf("username or password not found")
		return
	}

	return
}
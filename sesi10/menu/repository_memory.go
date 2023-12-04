package menu

import "database/sql"

/*
Repository ini berfungsi sebagai temporary
atau latihan sebelum kita masuk ke materi database
*/
type MemoryRepo struct{}

var data = []Menu{}

func NewMemoryRepo() MemoryRepo {
	return MemoryRepo{}
}

// GetAll implements RepositoryInterface
func (m MemoryRepo) GetAll() (menus []Menu, err error) {
	menus = data
	return
}

// GetById implements RepositoryInterface
func (m MemoryRepo) GetById(id int) (menu Menu, err error) {
	for _, d := range data {
		if d.Id == id {
			return d, nil
		}
	}

	err = sql.ErrNoRows
	return
}

// Save implements RepositoryInterface
func (m MemoryRepo) Save(menu Menu) (err error) {
	// menu.Id = len(data) + 1 // untuk automatic ID
	data = append(data, menu)
	return
}
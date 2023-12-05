package menu

import (
	"database/sql"
	// "fmt"
)

/*
membuat struct repository
struct ini mempunyai dependency ke *sql.DB
yang mana merupakan sebuah object koneksi database
*/
type Repository struct {
	db *sql.DB
}

/*
function untuk inisiasi repository dengan memasukkan dependency
ke dalam parameter.
*/
func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}



func (r Repository) GetAll() (menus []Menu, err error) {
	// query database
	query := `
		SELECT 
			*
		FROM
			menus
	`

	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// proses untuk mengeksekusi query ke database
	rows, err := stmt.Query()
	// fmt.Print(rows)
	if err != nil {
		return
	}

	// close rows
	defer rows.Close()

	for rows.Next() {
		menu := Menu{}
		// urutan scan-nya harus sama dengan
		// saat melakukan query
		// dan selalu gunakan kirim alamat memory (&)
		// jika ingin memasukkan suatu nilai
		err = rows.Scan(
			&menu.Id, &menu.Name, &menu.Category,
			&menu.Desc, &menu.Price,&menu.ImageUrl,
			&menu.CreatedAt,&menu.UpdatedAt,
		)
		if err != nil {
			return
		}
		// proses menambahakan ke dalam menus yang akan menjadi return value
		menus = append(menus, menu)

	}

	return

}

func (r Repository) GetById(id int) (menu Menu, err error) {
	// query database
	query := `
		SELECT 
			*
		FROM
			menus
		WHERE id = $1
	`

	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// eksekusi query, namun yang diambil adalah 1 baris.
	// dan memasukkan argument berupa `id` karena pada query
	// kita mempunyai parameter ($1).
	row := stmt.QueryRow(id)

	// proses mengambil value dari database
	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Category,
		&menu.Desc, &menu.Price,&menu.ImageUrl,
		&menu.CreatedAt, &menu.UpdatedAt,
	)

	return
}

func (r Repository) Save(menu Menu) (err error) {
	// query database
	query := `
		INSERT INTO menus (name, category, description, price, image_url ,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// eksekusi query, sesuai dengan urutan dari parameter nya
	_, err = stmt.Exec(menu.Name, menu.Category, menu.Desc, menu.Price, menu.ImageUrl,menu.CreatedAt, menu.UpdatedAt)
	return
}


func (r Repository) UpdateById(id int, menu Menu) (err error) {
	// query database
	query := `
		UPDATE menus
		SET 
		name = $1,
		category = $2,
		description = $3,
		price = $4,
		image_url = $5,
		updated_at = $6
		WHERE
		id = $7
	`
	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// eksekusi query, sesuai dengan urutan dari parameter nya
	_, err = stmt.Exec(menu.Name, menu.Category, menu.Desc, menu.Price, menu.ImageUrl,menu.UpdatedAt, id)
	return
}
package employee

import (
	"database/sql"
	"fmt"
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

func (r Repository) Save(employee Employee) (err error) {
	// query database
	query := `
		INSERT INTO employees (name, address, nip, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// eksekusi query, sesuai dengan urutan dari parameter nya
	_, err = stmt.Exec(employee.Name, employee.Address,employee.Nip,  employee.CreatedAt, employee.UpdatedAt)
	return
}


func (r Repository) GetById(id int) (employee Employee, err error) {
	// query database
	query := `
		SELECT 
			*
		FROM
			employees
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
		&employee.Id,&employee.Name,&employee.Nip,&employee.Address,&employee.CreatedAt,&employee.UpdatedAt,
	)

	return
}

func (r Repository) GetAll() (menus []Employee, err error) {
	// query database
	query := `
		SELECT 
		id, name, nip
		, address
		, created_at, updated_at
		FROM
			employees
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
	fmt.Print(rows)
	if err != nil {
		return
	}

	// close rows
	defer rows.Close()

	for rows.Next() {
		menu := Employee{}
		// urutan scan-nya harus sama dengan
		// saat melakukan query
		// dan selalu gunakan kirim alamat memory (&)
		// jika ingin memasukkan suatu nilai
		err = rows.Scan(
			&menu.Id, &menu.Name, &menu.Nip,
			&menu.Address,
			&menu.CreatedAt, &menu.UpdatedAt,
		)
		if err != nil {
			return
		}
		// proses menambahakan ke dalam menus yang akan menjadi return value
		menus = append(menus, menu)

	}

	return

}

func (r Repository) DeleteById(id int) (err error) {
	// query database
	query := `
		DELETE FROM employees
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

	// proses mengambil value dari database
	stmt.Exec(id)
	return
}
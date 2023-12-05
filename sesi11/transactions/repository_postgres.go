package transactions

import (
	"database/sql"
	// "fmt"
)



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

func (r Repository) Save(data Transactions) (err error) {
	// query database
	query := `
		INSERT INTO transactions (employee_id,menu_id,quantity,total,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	// set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	// selalu close stmt
	defer stmt.Close()

	// eksekusi query, sesuai dengan urutan dari parameter nya
	_, err = stmt.Exec(data.EmployeeId,data.MenuId,data.Quantity,data.Total,data.CreatedAt,data.UpdatedAt)
	return
}

func (r Repository) GetEmployeeById(id int) (emp Employee, err error) {
	query := `SELECT id, name FROM employees WHERE id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&emp.Id, &emp.Name,
	)
	return
}

// GetMenuById implements RepositoryInterface
func (r Repository) GetMenuById(id int) (menu Menu, err error) {
	query := `SELECT id, name, price FROM menus where id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Price,
	)
	return
}

func (r Repository) GetAll() (data []Transactions, err error) {
	// query database
	query := `
		SELECT 
			*
		FROM
			transactions
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
		datas := Transactions{}
		// urutan scan-nya harus sama dengan
		// saat melakukan query
		// dan selalu gunakan kirim alamat memory (&)
		// jika ingin memasukkan suatu nilai
		err = rows.Scan(&datas.Id,
			&datas.EmployeeId,&datas.MenuId,&datas.Quantity,&datas.Total,&datas.CreatedAt,&datas.UpdatedAt,
		)
		if err != nil {
			return
		}
		// // proses menambahakan ke dalam menus yang akan menjadi return value
		data = append(data, datas)

	}

	return

}


func (r Repository) FindById(id int) (data Transactions, err error) {
	query := `SELECT t.id, t.employee_id, t.menu_id, t.quantity, t.total, t.created_at, t.updated_at, 
				m.id, m.name, m.price, 
				e.id, e.name
				FROM transactions as t
				JOIN employees as e
					ON e.id = t.employee_id
				JOIN menus as m
					ON m.id = t.menu_id
				WHERE t.id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&data.Id, &data.EmployeeId, &data.MenuId, &data.Quantity, &data.Total, &data.CreatedAt, &data.UpdatedAt,
		&data.Menu.Id, &data.Menu.Name, &data.Menu.Price,
		&data.Employee.Id, &data.Employee.Name,
	)
	return
}
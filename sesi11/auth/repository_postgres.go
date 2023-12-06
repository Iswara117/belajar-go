package auth

import "database/sql"

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

// object repository
// mempunyai depedensi ke object db 
type Repository struct {
	db *sql.DB
}

func (r Repository) Create(auth Auth) (err error) {
	// query database
	query := `
		INSERT INTO auth (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(auth.Email, auth.Password, auth.CreatedAt, auth.UpdatedAt)
	return
}

// GetByEmail implements RepositoryInterface
func (r Repository) GetByEmail(email string) (auth Auth, err error) {
	query := `
		SELECT 
			id, email, password
			, created_at, updated_at
		FROM auth
		WHERE email = $1
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)

	err = row.Scan(
		&auth.Id, &auth.Email, &auth.Password,
		&auth.CreatedAt, &auth.UpdatedAt,
	)

	return
}

func (r Repository) GetById(id interface{}) (auth Auth, err error) {
	// query database
	query := `
		SELECT 
			*
		FROM
			auth
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
		&auth.Id,&auth.Email,&auth.Password,&auth.UpdatedAt,&auth.CreatedAt,
	)

	return
}

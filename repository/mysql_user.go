package repository

import (
	"database/sql"
	"errors"

	"github.com/ivantedja/tuku"

	"github.com/jmoiron/sqlx"
)

type mysqlUserRepo struct {
	db *sqlx.DB
}

func NewMysqlUserRepo(db *sql.DB) tuku.UserRepo {
	sqlxdb := sqlx.NewDb(db, "mysql")
	return &mysqlUserRepo{sqlxdb}
}

func (ur *mysqlUserRepo) Get(id int64) (*tuku.User, error) {
	query := `SELECT id, name, admin, updated_at, created_at FROM users WHERE id = ?`
	var user tuku.User
	err := ur.db.Get(&user, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *mysqlUserRepo) Create(user *tuku.User) error {
	query := `INSERT users SET name=? , admin=?, updated_at=? , created_at=?`
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Name, user.Admin, user.UpdatedAt, user.CreatedAt)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = lastID
	return nil
}

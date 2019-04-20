package repository

import (
	"database/sql"
	"errors"

	"github.com/ivantedja/tuku"

	"github.com/jmoiron/sqlx"
)

type mysqlProductRepo struct {
	db *sqlx.DB
}

func NewMysqlProductRepo(db *sql.DB) tuku.ProductRepo {
	sqlxdb := sqlx.NewDb(db, "mysql")
	return &mysqlProductRepo{sqlxdb}
}

func (pr *mysqlProductRepo) Create(p *tuku.Product) error {
	query := `INSERT products SET user_id=?, name=?, price=?, quantity=?, updated_at=? , created_at=?`
	stmt, err := pr.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.UserID, p.Name, p.Price, p.Quantity, p.UpdatedAt, p.CreatedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = lastID
	return nil
}

func (p *mysqlProductRepo) GetProductsByUserID(limit int64, offset int64, userID int64) ([]tuku.Product, int64, error) {
	if limit == 0 {
		limit = 10
	}

	if userID == 0 {
		return nil, 0, errors.New("No user provided")
	}

	var count int
	var results []tuku.Product

	err := p.db.Get(&count, "SELECT COUNT(id) FROM products WHERE user_id = ?", userID)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, user_id, name, price, quantity, updated_at, created_at FROM products WHERE user_id = ? LIMIT ?,? `

	err = p.db.Select(&results, query, userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return results, int64(count), err
}

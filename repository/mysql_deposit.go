package repository

import (
	"database/sql"
	"errors"

	"github.com/ivantedja/tuku"

	"github.com/jmoiron/sqlx"
)

type mysqlDepositRepo struct {
	db *sqlx.DB
}

func NewMysqlDepositRepo(db *sql.DB) tuku.DepositRepo {
	sqlxdb := sqlx.NewDb(db, "mysql")
	return &mysqlDepositRepo{sqlxdb}
}

func (d *mysqlDepositRepo) GetBalanceByUserID(userID int64) (*tuku.Deposit, error) {
	query := `SELECT id, user_id, price, updated_at, created_at FROM deposits WHERE user_id = ?`
	var deposit tuku.Deposit
	err := d.db.Get(&deposit, query, userID)
	if err == sql.ErrNoRows {
		return nil, errors.New("Deposit not found")
	}

	if err != nil {
		return nil, err
	}

	return &deposit, nil
}

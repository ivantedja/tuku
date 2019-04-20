package repository

import (
	"database/sql"
	"errors"
	"fmt"

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

func (dr *mysqlDepositRepo) Get(ID int64) (*tuku.Deposit, error) {
	query := `SELECT id, user_id, balance, updated_at, created_at FROM deposits WHERE id = ?`
	var deposit tuku.Deposit
	err := dr.db.Get(&deposit, query, ID)
	if err == sql.ErrNoRows {
		return nil, errors.New("Deposit not found")
	}

	if err != nil {
		return nil, err
	}

	return &deposit, nil
}


func (dr *mysqlDepositRepo) Update(ID int64, d *tuku.Deposit) error {
	query := `UPDATE deposits set balance=?, updated_at=? WHERE id = ?`

	stmt, err := dr.db.Prepare(query)
	if err != nil {
		return nil
	}

	res, err := stmt.Exec(d.Balance, d.UpdatedAt, ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affected)
		return err
	}

	return nil
}

func (d *mysqlDepositRepo) GetBalanceByUserID(userID int64) (*tuku.Deposit, error) {
	query := `SELECT id, user_id, balance, updated_at, created_at FROM deposits WHERE user_id = ?`
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

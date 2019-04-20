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

func (d *mysqlDepositRepo) Get(ID int64) (*tuku.Deposit, error) {
	query := `SELECT id, user_id, balance, updated_at, created_at FROM deposits WHERE id = ?`
	var deposit tuku.Deposit
	err := d.db.Get(&deposit, query, ID)
	if err == sql.ErrNoRows {
		return nil, errors.New("Deposit not found")
	}

	if err != nil {
		return nil, err
	}

	return &deposit, nil
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

func (d *mysqlDepositRepo) UpdateDeposit(ID int64, deposit *tuku.Deposit) error {
	query := `UPDATE deposits set balance=?, updated_at=? WHERE id = ?`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil
	}

	res, err := stmt.Exec(deposit.Balance, deposit.UpdatedAt, ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}

	return nil
}

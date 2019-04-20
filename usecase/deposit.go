package usecase

import (
	"github.com/ivantedja/tuku"
)

type depositUsecase struct {
	DepositRepo tuku.DepositRepo
}

func NewDepositUsecase(dr tuku.DepositRepo) tuku.DepositUsecase {
	return &depositUsecase{dr}
}

func (du *depositUsecase) Get(ID int64) (*tuku.Deposit, error) {
	return du.DepositRepo.Get(ID)
}

func (du *depositUsecase) GetBalanceByUserID(userID int64) (*tuku.Deposit, error) {
	return du.DepositRepo.GetBalanceByUserID(userID)
}

func (du *depositUsecase) ReduceBalance(ID int64, amount int64) error {
	deposit, err := du.Get(ID)
	if err != nil {
		return err
	}

	// TODO: cover negative balance via unit test
	deposit.Balance -= amount

	err = du.DepositRepo.UpdateDeposit(ID, deposit)
	return err
}

func (du *depositUsecase) IncreaseBalance(ID int64, amount int64) error {
	deposit, err := du.Get(ID)
	if err != nil {
		return err
	}

	deposit.Balance += amount

	err = du.DepositRepo.UpdateDeposit(ID, deposit)
	return err
}

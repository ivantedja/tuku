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

	err = du.DepositRepo.UpdateDeposit(deposit.ID, deposit)
	return err
}

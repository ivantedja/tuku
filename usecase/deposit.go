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

func (du *depositUsecase) GetBalanceByUserID(userID int64) (*tuku.Deposit, error) {
	return du.DepositRepo.GetBalanceByUserID(userID)
}

package usecase

import (
	"errors"
	"fmt"

	"github.com/ivantedja/tuku"
)

type depositUsecase struct {
	DepositRepo tuku.DepositRepo
	ProductRepo tuku.ProductRepo
}

func NewDepositUsecase(dr tuku.DepositRepo, pr tuku.ProductRepo) tuku.DepositUsecase {
	return &depositUsecase{dr, pr}
}

func (du *depositUsecase) GetBalanceByUserID(userID int64) (*tuku.Deposit, error) {
	return du.DepositRepo.GetBalanceByUserID(userID)
}

func (du *depositUsecase) ReduceBalance(userID int64, productID int64, quantity int64) error {
	deposit, err := du.GetBalanceByUserID(userID)
	if err != nil {
		return err
	}

	product, err := du.ProductRepo.Get(productID)
	if err != nil {
		return err
	}

	if product.Quantity < quantity {
		return errors.New("Quantity not enough")
	}

	//allow negative value
	//if deposit.Balance < quantity * product.Price {
	//	return errors.New("Balance not enough")
	//}

	deposit.Balance -= product.Price * quantity
	err = du.DepositRepo.UpdateDeposit(userID, deposit)

	return err
}

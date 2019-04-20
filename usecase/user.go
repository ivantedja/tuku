package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/ivantedja/tuku"
)

type userUsecase struct {
	DepositUsecase tuku.DepositUsecase
	UserRepo       tuku.UserRepo
	ProductRepo    tuku.ProductRepo
}

func NewUserUsecase(du tuku.DepositUsecase, ur tuku.UserRepo, pr tuku.ProductRepo) tuku.UserUsecase {
	return &userUsecase{du, ur, pr}
}

func (u *userUsecase) GetUser(id int64) (*tuku.User, error) {
	return u.UserRepo.Get(id)
}

func (u *userUsecase) CreateUser(user *tuku.User) error {
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	return u.UserRepo.Create(user)
}

func (uu *userUsecase) Pay(userID int64, productID int64) error {
	balance, _ := uu.DepositUsecase.GetBalanceByUserID(userID)
	fmt.Printf("pay balance: %+v\n", balance)

	product, count, _ := uu.ProductRepo.GetProductsByUserID(10, 0, userID)
	fmt.Printf("pay count: %d, product: %+v\n", count, product)

	return nil
}

func (uu *userUsecase) CreateProduct(userID int64, p *tuku.Product) error {
	now := time.Now().UTC()
	p.UserID = userID
	p.CreatedAt = now
	p.UpdatedAt = now

	return uu.ProductRepo.Create(p)
}

func (uu *userUsecase) GetProductsByUserID(limit int64, offset int64, userID int64) ([]tuku.Product, int64, error) {
	if limit == 0 {
		limit = 10
	}

	if userID == 0 {
		return nil, 0, errors.New("No user provided")
	}

	return uu.ProductRepo.GetProductsByUserID(limit, offset, userID)
}

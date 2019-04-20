package tuku

type UserUsecase interface {
	GetUser(id int64) (*User, error)
	CreateUser(user *User) error

	Pay(userID int64, productID int64, quantity int64) error

	CreateProduct(userID int64, product *Product) error
	GetProductsByUserID(limit int64, offset int64, userID int64) ([]Product, int64, error)
}

type DepositUsecase interface {
	Get(ID int64) (*Deposit, error)
	GetBalanceByUserID(userID int64) (*Deposit, error)
	ReduceBalance(ID int64, amount int64) error
	IncreaseBalance(ID int64, amount int64) error
}

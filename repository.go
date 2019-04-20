package tuku

type UserRepo interface {
	Get(id int64) (*User, error)
	Create(user *User) error
}

type ProductRepo interface {
	Get(id int64) (*Product, error)
	Create(product *Product) error
	Update(ID int64, product *Product) error
	GetProductsByUserID(limit int64, offset int64, userID int64) ([]Product, int64, error)
}

type DepositRepo interface {
	Get(ID int64) (*Deposit, error)
	Update(ID int64, deposit *Deposit) error
	GetBalanceByUserID(userID int64) (*Deposit, error)
}

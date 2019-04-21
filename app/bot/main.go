package main

import (
	"database/sql"
	"fmt"

	"github.com/ivantedja/tuku/config"
	"github.com/ivantedja/tuku/delivery"
	"github.com/ivantedja/tuku/repository"
	"github.com/ivantedja/tuku/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Charset))
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.Database.Pool)
	err = db.Ping()
	return db, err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cfg, err := config.NewConfig()
	check(err)

	db, err := NewMysql(cfg)
	check(err)

	userRepo := repository.NewMysqlUserRepo(db)
	productRepo := repository.NewMysqlProductRepo(db)
	depositRepo := repository.NewMysqlDepositRepo(db)

	depositUsecase := usecase.NewDepositUsecase(depositRepo)
	userUsecase := usecase.NewUserUsecase(depositUsecase, userRepo, productRepo)

	slackRtm := delivery.NewTukuSlackRtm(cfg, userUsecase, depositUsecase)
	err = slackRtm.ListenAndServe()
	fmt.Printf("Slack RTM error: %+v", err)
}

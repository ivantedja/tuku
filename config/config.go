package config

import (
	"database/sql"
	"fmt"
	"github.com/joeshaw/envdecode"

	"github.com/subosito/gotenv"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	gotenv.Load(".env")
}

//Config represent global config
type Config struct {
	Database struct {
		Username string `env:"DATABASE_USERNAME,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		Host     string `env:"DATABASE_HOST,default=localhost"`
		Port     string `env:"DATABASE_PORT,default=3306"`
		Name     string `env:"DATABASE_NAME,required"`
		Charset  string `env:"DATABASE_CHARSET,default=utf8"`
		Pool     int    `env:"DATABASE_POOL,default=50"`
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewMysql() (*sql.DB, error) {
	cfg := &Config{}
	err := envdecode.Decode(cfg)
	check(err)

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

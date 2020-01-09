package db

//MySQL连接初始化
import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)


//初始化
func Init(dsn string) error {
	database, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open mysql error", err)
		return errors.New("open mysql error")
	}
	DB = database

	//测试连接成功与否
	err = DB.Ping()
	if err != nil {
		fmt.Println("connect error", err)
		return errors.New("ping error")
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	fmt.Println("连接DB成功")
	return nil
}
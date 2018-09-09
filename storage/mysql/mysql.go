package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
)

type mysql struct {
	DB *sql.DB
}

const tablename = "todos"

var (
	user   = os.Getenv("MYSQL_USER")
	pass   = os.Getenv("MYSQL_PASS")
	dbname = os.Getenv("MYSQL_DATABASE")
)

func New() (*mysql, error) {
	connection := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
	db, err := xorm.NewEngine("mysql", connection)
	if err != nil {
		return nil, err
	}

	return &mysql{
		DB: db,
	}, nil
}

func Save(long_url string) error {

	return nil
}

func CountUrl(code string) error {

	return nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

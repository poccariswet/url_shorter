package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"github.com/soeyusuke/url_shorter/storage"
)

type mysql struct {
	DB *xorm.Engine
}

const tablename = "urlsho"

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

func (m *mysql) Save(long_url string) error {
	affected, err := m.DB.Insert(&storage.Urlsho{LongURL: long_url})
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf(": %d", affected))
	}

	return nil
}

func (m *mysql) CountUrl(code string) error {

	return nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

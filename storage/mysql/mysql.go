package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/soeyusuke/gitclone/gorm"
	"github.com/soeyusuke/url_shorter/base62"
	"github.com/soeyusuke/url_shorter/storage"
)

type mysql struct {
	DB *gorm.DB
}

const (
	user      = "urlsho_user"
	pass      = "urlsho_pass"
	dbname    = "urlsho_db"
	tablename = "urlsho"
)

func Init() *mysql {
	return &mysql{}
}

func (m *mysql) NewDB() error {
	connection := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return err
	}
	m.DB = db

	return nil
}

func (m *mysql) Save(long_url string) (string, error) {
	if err := m.NewDB(); err != nil {
		return "", err
	}
	defer m.Close()

	u := storage.Urlsho{
		LongURL: long_url,
	}
	m.DB.Table(tablename).Create(u)

	return base62.Encode(u.Id), nil
}

func (m *mysql) CountUrl(code string) error {
	if err := m.NewDB(); err != nil {
		return err
	}
	defer m.Close()

	//TODO m.DB.Update()

	return nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/soeyusuke/gitclone/gorm"
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

func InitConfig() *mysql {
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

func (m *mysql) Save(long_url string) error {
	if err := m.NewDB(); err != nil {
		return err
	}
	defer m.Close()

	m.DB.Create(&storage.Urlsho{
		LongURL: long_url,
	})

	return nil
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

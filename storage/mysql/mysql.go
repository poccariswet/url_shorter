package mysql

import (
	"errors"
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
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
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

	u := &storage.Urlsho{
		LongURL: long_url,
		Count:   0,
	}
	m.DB.Table(tablename).Create(u)
	if u.Id == 0 {
		return "", errors.New("not save data")
	}

	return base62.Encode(u.Id), nil
}

func (m *mysql) LoadAndCountUp(id int) (string, error) {
	if err := m.NewDB(); err != nil {
		return "", err
	}
	defer m.Close()

	var u storage.Urlsho
	m.DB.Raw("UPDATE urlsho SET count=count+1 WHERE id = ?", id)
	m.DB.Table(tablename).Find(&u, "id = ?", id)
	if u.Id == 0 {
		return "", errors.New(fmt.Sprintf("%v", u))
	}

	return u.LongURL, nil
}

func (m *mysql) FetchInfo(id int) (*storage.Urlsho, error) {
	if err := m.NewDB(); err != nil {
		return nil, err
	}
	defer m.Close()

	var u storage.Urlsho
	m.DB.Table(tablename).Find(&u, "id = ?", id)
	if u.Id == 0 {
		return nil, errors.New("fetch data err")
	}

	return &u, nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

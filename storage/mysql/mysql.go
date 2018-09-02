package mysql

import (
	"database/sql"

	"github.com/soeyusuke/gitclone/ursho/storage"
)

type mysql struct {
	DB *sql.DB
}

func New() (storage.Service, error) {

	return storage.Service{}, nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

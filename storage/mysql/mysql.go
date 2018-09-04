package mysql

import (
	"database/sql"
)

type mysql struct {
	DB *sql.DB
}

func New() (*mysql, error) {
	return &mysql{}, nil
}

func (m *mysql) Close() error {
	return m.DB.Close()
}

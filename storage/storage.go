package storage

type Service interface {
	Close() error
}

type Shorten struct{}

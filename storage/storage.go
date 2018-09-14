package storage

type Service interface {
	Save(string) error
	CountUrl(string) error
	Load() (*Urlsho, error)
	Close() error
}

type Urlsho struct {
	Id      int    `json:"id"`
	LongURL string `json:"long_url"`
	Count   int    `json:"count"`
}

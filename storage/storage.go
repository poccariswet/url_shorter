package storage

type Service interface {
	Save(string) error
	CountUrl(string) error
	Close() error
}

type Urlsho struct {
	LongURL    string `json"long_url"`
	ShortenURL string `json:"short_url"`
	Count      int    `json:"count"`
}

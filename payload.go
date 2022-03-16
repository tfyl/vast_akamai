package vast_akamai

import "errors"

type Payload struct {
	Abck      string `json:"abck"`
	Url       string `json:"url"`
	UserAgent string `json:"userAgent"`
}

var (
	ErrMissingAbck = errors.New("missing abck")
	ErrMissingUrl  = errors.New("missing url")
)

func (p *Payload) validate() error {
	// Initial invalid _abck from akamai
	if p.Abck == "" {
		return ErrMissingAbck
	}

	// Url of site
	// example : https://www.nike.com/gb/
	if p.Url == "" {
		return ErrMissingUrl
	}

	return nil
}

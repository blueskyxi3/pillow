package ub

import (
	"net/url"
	"path/filepath"

	"github.com/enenumxela/to/pkg/to"
)

// URLBuilder
type URLBuilder struct {
	URL *url.URL
}

// NewURLBuilder
func NewURLBuilder(URL string) (ub *URLBuilder) {
	parsedURL, _ := url.Parse(URL)

	ub = &URLBuilder{
		URL: parsedURL,
	}

	return
}

// AddPath
func (ub *URLBuilder) AddPath(paths ...string) *URLBuilder {
	for _, path := range paths {
		ub.URL.Path = filepath.Join(ub.URL.Path, path)
	}

	return ub
}

// AddQuery
func (ub *URLBuilder) AddQuery(parameters map[string]interface{}) *URLBuilder {
	query := ub.URL.Query()

	for k, v := range parameters {
		query.Set(k, to.String(v))
	}

	ub.URL.RawQuery = query.Encode()

	return ub
}

// String
func (ub *URLBuilder) String() (path string) {
	return ub.URL.String()
}

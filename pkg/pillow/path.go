package pillow

import (
	"net/url"
	"path/filepath"

	"github.com/enenumxela/to/pkg/to"
)

type PathBuilder struct {
	URL *url.URL
}

func NewPathBuilder(URL string) (builder *PathBuilder) {
	parsedURL, _ := url.Parse(URL)

	builder = &PathBuilder{
		URL: parsedURL,
	}

	return
}

func (builder *PathBuilder) AddPath(paths ...string) *PathBuilder {
	for _, path := range paths {
		builder.URL.Path = filepath.Join(builder.URL.Path, path)
	}

	return builder
}

func (builder *PathBuilder) AddQuery(options Options) *PathBuilder {
	query := builder.URL.Query()

	for k, v := range options {
		query.Set(k, to.String(v))
	}

	builder.URL.RawQuery = query.Encode()

	return builder
}

func (builder *PathBuilder) String() (path string) {
	return builder.URL.String()
}

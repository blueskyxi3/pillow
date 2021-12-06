package pillow

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/enenumxela/pillow/pkg/ub"
	"github.com/imdario/mergo"
)

// Client
type Client struct {
	dsn  string
	http *http.Client
}

// ClientOptions
type ClientOptions struct {
	Timeout int
}

// New
func New(dsn string, opts ...*ClientOptions) (*Client, error) {
	options := &ClientOptions{
		Timeout: 30,
	}

	if len(opts) > 0 {
		if err := mergo.Merge(&options, opts[0], mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	parsedDSN, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	HTTPClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(options.Timeout) * time.Second,
				KeepAlive: time.Second,
				DualStack: true,
			}).DialContext,
		},
	}

	res, err := HTTPClient.Get(parsedDSN.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Non-OK HTTP status")
	}

	return &Client{
		dsn:  parsedDSN.String(),
		http: HTTPClient,
	}, nil
}

// DSN returns the data source name used to connect this client.
func (client *Client) DSN() string {
	return client.dsn
}

// Ping
func (client *Client) Ping(ctx context.Context) (pong bool, err error) {
	path := client.DSN()

	res, err := client.request(http.MethodHead, path, nil, nil)
	if err == nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		pong = true
	}

	return
}

// ListDatabases
func (client *Client) ListDatabases(ctx context.Context, options ...map[string]interface{}) (databases []string, err error) {
	path := ub.NewURLBuilder(client.DSN()).AddPath("_all_dbs").String()

	res, err := client.request(http.MethodGet, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&databases); err != nil {
		return
	}

	return
}

// Database
func (client *Client) Database(ctx context.Context, name string) (db *DB) {
	db = &DB{
		name,
		client,
	}

	return
}

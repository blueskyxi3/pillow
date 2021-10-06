package pillow

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/enenumxela/pillow/pkg/ub"
)

type Client struct {
	dsn  string
	http *http.Client
}

// New
func New(dsn string) (client *Client, err error) {
	parsedDSN, err := url.Parse(dsn)
	if err != nil {
		return
	}

	timeout := 30

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(timeout) * time.Second,
			KeepAlive: time.Second,
			DualStack: true,
		}).DialContext,
	}

	client = &Client{
		dsn: parsedDSN.String(),
		http: &http.Client{
			Transport: transport,
		},
	}

	return
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

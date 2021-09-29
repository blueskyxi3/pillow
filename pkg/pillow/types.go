package pillow

import "net/http"

type Client struct {
	dsn  string
	http *http.Client
}

// Options is a collection of options. The keys and values are backend specific.
type Options map[string]interface{}

type DB struct {
	name   string
	client *Client
}

type CreateDatabaseResponse struct {
	OK     bool   `json:"ok,omitempty"`
	Error  string `json:"error,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type DeleteDatabaseResponse struct {
	OK     bool   `json:"ok,omitempty"`
	Error  string `json:"error,omitempty"`
	Reason string `json:"reason,omitempty"`
}

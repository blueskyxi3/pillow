package pillow

import (
	"context"
	"encoding/json"
	"github.com/blueskyxi3/pillow/pkg/ub"
	"io/ioutil"
	"net/http"
)

type DB struct {
	name   string
	client *Client
}

func (db *DB) Name() string {
	return db.name
}

func (db *DB) Client() *Client {
	return db.client
}

func (db *DB) Exists(ctx context.Context, options ...map[string]interface{}) (exists bool, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name()).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodHead, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		exists = true
	}

	return
}

// Create
// {"ok":true}
// {"error":"file_exists","reason":"The database could not be created, the file already exists."}
func (db *DB) Create(ctx context.Context, options ...map[string]interface{}) (output *CreateDatabaseResponse, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name()).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodPut, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &output); err != nil {
		return
	}

	return
}

func (db *DB) Query(ctx context.Context, ddcoc, view string, options ...map[string]interface{}) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.name, "_design", ddcoc, "_view", view).AddQuery(mergeOptions(options...)).String()

	headers := map[string]string{
		"Accept": "application/json",
	}

	res, err := db.client.request(http.MethodGet, path, headers, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

//Delete {"ok":true}
func (db *DB) Delete(ctx context.Context, options ...map[string]interface{}) (output *DeleteDatabaseResponse, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name()).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodDelete, path, nil, nil)
	if err == nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &output); err != nil {
		return
	}

	return
}

package pillow

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/enenumxela/pillow/pkg/ub"
	"github.com/enenumxela/to/pkg/to"
)

// CheckDocument
func (db *DB) CheckDocument(ctx context.Context, id string, options ...Options) (exists bool, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), id).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodHead, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	output := map[string]interface{}{}

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	exists = true

	return
}

// CreateDocument
func (db *DB) CreateDocument(ctx context.Context, document interface{}, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name()).AddQuery(mergeOptions(options...)).String()

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	buf, err := json.Marshal(document)
	if err != nil {
		return
	}

	body := bytes.NewReader(buf)

	headers["Content-Length"] = to.String(len(buf))

	res, err := db.client.request(http.MethodPost, path, headers, body)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

func (db *DB) CreateDesignDocument(ctx context.Context, document map[string]interface{}, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), "_design", db.Name()).AddQuery(mergeOptions(options...)).String()

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	buf, err := json.Marshal(document)
	if err != nil {
		return
	}

	body := bytes.NewReader(buf)

	headers["Content-Length"] = to.String(len(buf))

	res, err := db.client.request(http.MethodPut, path, headers, body)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

// RetrieveDocument
func (db *DB) RetrieveDocument(ctx context.Context, id string, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), id).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodGet, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

// UpdateDocument
func (db *DB) UpdateDocument(ctx context.Context, id string, document interface{}, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), id).AddQuery(mergeOptions(options...)).String()

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	buf, err := json.Marshal(document)
	if err != nil {
		return
	}

	body := bytes.NewReader(buf)

	headers["Content-Length"] = to.String(len(buf))

	res, err := db.client.request(http.MethodPut, path, headers, body)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

// DeleteDocument
func (db *DB) DeleteDocument(ctx context.Context, id string, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), id).AddQuery(mergeOptions(options...)).String()

	res, err := db.client.request(http.MethodDelete, path, nil, nil)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		return
	}

	return
}

// ListDocuments
func (db *DB) ListDocuments(ctx context.Context, options ...Options) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), "_all_docs").AddQuery(mergeOptions(options...)).String()

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

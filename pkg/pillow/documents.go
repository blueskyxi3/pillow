package pillow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/blueskyxi3/pillow/pkg/ub"
	"github.com/enenumxela/to/pkg/to"
	"log"
	"net/http"
)

// CheckDocument check doc
func (db *DB) CheckDocument(ctx context.Context, id string, options ...map[string]interface{}) (exists bool, err error) {
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

// CreateDocument create doc
func (db *DB) CreateDocument(ctx context.Context, document interface{}, options ...map[string]interface{}) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name()).AddQuery(mergeOptions(options...)).String()
	log.Println("path:", path)
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

func (db *DB) CreateDesignDocument(ctx context.Context, document map[string]interface{}, options ...map[string]interface{}) (output map[string]interface{}, err error) {
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

// RetrieveDocument retrieve doc
func (db *DB) RetrieveDocument(ctx context.Context, id string, options ...map[string]interface{}) (output map[string]interface{}, err error) {
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

// UpdateDocument update doc
func (db *DB) UpdateDocument(ctx context.Context, id string, document interface{}, options ...map[string]interface{}) (output map[string]interface{}, err error) {
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

// DeleteDocument delete doc
func (db *DB) DeleteDocument(ctx context.Context, id string, options ...map[string]interface{}) (output map[string]interface{}, err error) {
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

// ListDocuments list docs
func (db *DB) ListDocuments(ctx context.Context, options ...map[string]interface{}) (output map[string]interface{}, err error) {
	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), "_all_docs").AddQuery(mergeOptions(options...)).String()

	headers := map[string]string{
		"Accept": "application/json",
	}
	log.Println("path:", path)
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
func (db *DB) QueryWithJSON(ctx context.Context, query string) (output map[string]interface{}, err error) {
	queryMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(query), &queryMap)
	if err != nil {
		return nil, err
	}
	bytesData, err := json.Marshal(queryMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)

	path := ub.NewURLBuilder(db.client.DSN()).AddPath(db.Name(), "_find").String()
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	//setDefault(&req.Header, "Accept", "application/json")
	//setDefault(&req.Header, "Content-Type", "application/json")
	log.Println("path:", path)
	res, err := db.client.request(http.MethodPost, path, headers, reader)
	if err != nil {
		log.Printf("-1---->%v \n", err)

		return
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		log.Println("-2---->", err.Error())
		return
	}

	return

}

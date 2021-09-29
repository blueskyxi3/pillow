package pillow

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (client *Client) request(method, path string, headers map[string]string, body io.Reader) (res *http.Response, err error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err = client.http.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode >= 400 {
		if res.Request.Method != http.MethodHead {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatalln(err)
			}

			var output struct {
				Error  string `json:"error"`
				Reason string `json:"reason"`
			}

			if err = json.Unmarshal(body, &output); err != nil {
				log.Fatalln(err)
			}

			err = &Error{
				HTTPStatus: res.StatusCode,
				Message:    output.Reason,
			}

			return nil, err
		}

	}

	return
}

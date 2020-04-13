package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Method      string
	Url         string
	ApiKey      string
	Body        interface{}
	ContentType ContentType
}

type ContentType string

const (
	ContentType_APPLICATION_JSON ContentType = "application/json"
)

func New(conf Config) ([]byte, int, error) {

	var b []byte
	var err error

	switch conf.ContentType {
	case ContentType_APPLICATION_JSON:
		b, err = json.Marshal(conf.Body)
		if err != nil {
			return []byte{}, 0, err
		}
	}
	req, err := http.NewRequest(conf.Method, conf.Url, bytes.NewReader(b))
	if err != nil {
		return []byte{}, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.ApiKey))

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	return b, resp.StatusCode, nil
}

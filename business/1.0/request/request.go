package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Config struct {
	Method      string
	Url         string
	AccessToken string
	Sandbox     bool
	Body        interface{}
	ContentType ContentType
}

type ContentType string

const (
	ContentType_APPLICATION_FORM ContentType = "application/x-www-form-urlencoded"
	ContentType_APPLICATION_JSON ContentType = "application/json"
)

func New(conf Config) ([]byte, int, error) {

	var b []byte
	var err error

	switch conf.ContentType {
	case ContentType_APPLICATION_FORM:
		b = []byte(conf.Body.(url.Values).Encode())

	case ContentType_APPLICATION_JSON:
		b, err = json.Marshal(conf.Body)
		if err != nil {
			return []byte{}, 0, err
		}
	}

	if conf.Sandbox {
		conf.Url = fmt.Sprintf("%ssandbox-%s", conf.Url[:8], conf.Url[8:])
	}

	req, err := http.NewRequest(conf.Method, conf.Url, bytes.NewReader(b))
	if err != nil {
		return []byte{}, 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.AccessToken))

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

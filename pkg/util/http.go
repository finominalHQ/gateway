package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

func Request(method string, url string, payload any) (any, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(JsonStringify(payload)))
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	d := make(Json)
	return JsonParse(body, &d), nil
}

func Post(url string, payload any) (any, error) {
	return Request(http.MethodPost, url, payload)
}

func Get(url string) (any, error) {
	return Request(http.MethodGet, url, nil)
}

func Patch(url string, payload any) (any, error) {
	return Request(http.MethodPatch, url, payload)
}

func Put(url string, payload any) (any, error) {
	return Request(http.MethodPut, url, payload)
}

func Delete(url string) (any, error) {
	return Request(http.MethodDelete, url, nil)
}

package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

func Request(method string, url string, payload any) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(JsonParse(payload)))
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

	var d map[string]interface{}
	return JsonStringify(body, &d).(map[string]interface{}), nil
}

func Post(url string, payload any) (map[string]interface{}, error) {
	return Request(http.MethodPost, url, payload)
}

func Get(url string) (map[string]interface{}, error) {
	return Request(http.MethodGet, url, nil)
}

func Patch(url string, payload any) (map[string]interface{}, error) {
	return Request(http.MethodPatch, url, payload)
}

func Put(url string, payload any) (map[string]interface{}, error) {
	return Request(http.MethodPut, url, payload)
}

func Delete(url string) (map[string]interface{}, error) {
	return Request(http.MethodDelete, url, nil)
}

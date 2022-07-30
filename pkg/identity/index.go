package identity

import (
	"fmt"
	"gateway/pkg/cache"
	"gateway/pkg/util"
	"net/http"
	"strings"

	"github.com/gobuffalo/envy"
)

var (
	baseUrl string = envy.Get("IDENTITY_BASE_URL", "")
)

func AuthN(req *http.Request) bool {
	paths := strings.Split(req.URL.Path, "")

	service := ""
	if len(paths) > 0 {
		service = paths[0]
	}

	resource := ""
	if len(paths) > 1 {
		resource = paths[1]
	}

	action := ""
	if len(paths) > 2 {
		action = paths[2]
	}

	key := fmt.Sprintf("gateway:identity:authn:%s-%s-%s-%s-%s-%s", req.Method, req.Header, req.URL.Port(), service, resource, action)
	res := cache.Get(key)
	if res == nil {
		payload := map[string]string{
			"method":   req.Method,
			"host":     req.URL.Host,
			"port":     req.URL.Port(),
			"service":  service,
			"resource": resource,
			"action":   action,
		}

		var err error
		res, err = util.Post(baseUrl+"/user/authenticate", payload)
		if err != nil {
			return false
		}

		cache.Set(key, res, 60*10)
	}

	data := res.(map[string]any)["data"]
	status := data.(map[string]any)["status"]
	if status != true {
		return false
	}

	return true
}

func AuthZ(req *http.Request) bool {
	paths := strings.Split(req.URL.Path, "")

	service := ""
	if len(paths) > 0 {
		service = paths[0]
	}

	resource := ""
	if len(paths) > 1 {
		resource = paths[1]
	}

	action := ""
	if len(paths) > 2 {
		action = paths[2]
	}

	key := fmt.Sprintf("gateway:identity:authz:%s-%s-%s-%s-%s-%s", req.Method, req.Header, req.URL.Port(), service, resource, action)
	res := cache.Get(key)
	if res == nil {
		payload := map[string]string{
			"method":   req.Method,
			"host":     req.URL.Host,
			"port":     req.URL.Port(),
			"service":  service,
			"resource": resource,
			"action":   action,
		}

		var err error
		res, err = util.Post(baseUrl+"/user/authorize", payload)
		if err != nil {
			return false
		}

		cache.Set(key, res, 60*10)
	}

	data := res.(map[string]any)["data"]
	status := data.(map[string]any)["status"]
	if status != true {
		return false
	}

	return true
}

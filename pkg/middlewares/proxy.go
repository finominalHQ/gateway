package middlewares

import (
	"fmt"
	"gateway/pkg/cache"
	"gateway/pkg/route"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

// Proxy request to the right backend
// Proxy is a raw go middleware, NOT a buffalo middleware
// It is registered as a PreWare with the buffalo app
type Proxy struct {
	next http.Handler
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	key := fmt.Sprintf("gateway:route:%s-%s-%s-%s", r.Method, r.URL.Host, r.URL.Port(), strings.ReplaceAll(r.URL.Path, "/", "-"))
	var routes *route.Routes
	routes, _ = cache.Get(key).(*route.Routes)
	if routes == nil {
		slugs := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		var service string
		if len(slugs) > 0 {
			service = slugs[0]
		}

		routes = &route.Routes{}
		_ = route.GetRoutesByService(service, routes)
	}

	if routes == nil {
		p.next.ServeHTTP(w, r)
		return
	}

	var routeObj *route.Route
	for _, v := range *routes {
		match, _ := regexp.MatchString(v.Pattern, r.URL.String())

		if match {
			routeObj = &v
			cache.Set(key, routeObj, 60*60)
		}
	}

	// skip this middleware
	// do not proxy request if no matching outgoing route is found
	if routeObj == nil {
		p.next.ServeHTTP(w, r)
		return
	}

	rp := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// mutate req header
			// for k, v := range routeObj.Header {
			// 	req.Header.Set(k, v.(string))
			// }

			// mutate req query
			// existingQueries := req.URL.Query()
			// for k, v := range routeObj.Query {
			// 	existingQueries.Add(k, v.(string))
			// }
			// req.URL.RawQuery = existingQueries.Encode()

			// mutate req body
			// b, _ := ioutil.ReadAll(req.Body)
			// defer req.Body.Close()
			// existingBody := make(util.Json)
			// util.JsonParse(b, &existingBody)
			// for k, v := range routeObj.Body {
			// 	existingBody[k] = v
			// }

			// mutate/swap incoming host to upstream host
			parsedUrl, err := url.ParseRequestURI(routeObj.Upstream)
			if err != nil {
				// throw error
			}

			req.URL.Scheme = parsedUrl.Scheme
			req.URL.Host = parsedUrl.Host // host:port

			req.URL.Path = "/" + strings.Join(strings.Split(strings.Trim(r.URL.Path, "/"), "/")[1:], "/")
		},
	}

	rp.ServeHTTP(w, r)
}

func ProxyHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := Proxy{next: next}
		p.ServeHTTP(w, r)
	})
}

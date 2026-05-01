package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func PreviewProxy(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid preview path", http.StatusBadRequest)
		return
	}

	golet_id := parts[2]
	targetURL := "http://" + "preview." + golet_id + ".localhost:3006"
	target, _ := url.Parse(targetURL)

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(target)

			prefix := "/preview/" + golet_id
			pr.Out.URL.Path = strings.TrimPrefix(pr.In.URL.Path, prefix)

			if pr.Out.URL.Path == "" {
				pr.Out.URL.Path = "/"
			}
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Println("Proxy error:", err)
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
		},
	}
	proxy.ServeHTTP(w, r)

}

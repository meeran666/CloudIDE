package helpers

import (
	"frontend/models"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(golet_id string, targetURL string) {

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal("Invalid target URL:", err)
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {

			//  Set correct upstream URL
			pr.SetURL(target)

			// CRITICAL: fix Host header
			pr.Out.Host = target.Host

			// Optional debugging
			log.Println("Forwarding to:", pr.Out.URL.String())
		},

		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Println("Proxy error:", err)
			http.Error(w, "Backend unavailable", http.StatusBadGateway)
		},
	}

	models.Proxy = proxy
}

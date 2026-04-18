package helpers

import (
	"fmt"
	"frontend/models"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(golet_id string, targetURL string) {

	target, err := url.Parse(targetURL)
	fmt.Println(targetURL)
	if err != nil {
		log.Fatal("Invalid target URL:", err)
	}
	models.Proxy = httputil.NewSingleHostReverseProxy(target)

	models.Proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Println("Proxy error:", err)
		http.Error(w, "Backend unavailable", http.StatusBadGateway)
	}

}

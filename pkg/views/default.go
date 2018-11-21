package views

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Config struct {
	URL string `env:"VIEW_URL" envDefault:"http://localhost:3001"`
}

type container interface {
	ViewConfig() Config
}

func reverseProxy(target string) http.HandlerFunc {
	fmt.Println("Create view proxy at %s", target)
	url, _ := url.Parse(target)
	handler := httputil.NewSingleHostReverseProxy(url)
	return handler.ServeHTTP
}

func CreateViewsHandler(container container) http.HandlerFunc {
	return reverseProxy(container.ViewConfig().URL)
}

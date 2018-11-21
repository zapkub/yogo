package views

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Config struct {
	URL string `env:"VIEW_URL" envDefault:"http://localhost:3001"`
}

type container interface {
	ViewConfig() Config
}

func reverseProxy(target string) http.HandlerFunc {
	fmt.Printf("Create view proxy at %s\n", target)
	url, _ := url.Parse(target)
	handler := httputil.NewSingleHostReverseProxy(url)

	// line below is only for
	// development with nextjs hot reload
	// webpack hmr use Http event stream to
	// update hot loader status, if not provide this
	// flush interval, event stream will not flushing
	// message until request is close
	handler.FlushInterval = 100 * time.Millisecond
	return handler.ServeHTTP
}

func CreateViewsHandler(container container) http.HandlerFunc {
	return reverseProxy(container.ViewConfig().URL)
}

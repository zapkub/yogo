package views

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type container interface{}

func reverseProxy(target string) gin.HandlerFunc {

	return func(c *gin.Context) {
		url, _ := url.Parse(target)
		handler := httputil.NewSingleHostReverseProxy(url)
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

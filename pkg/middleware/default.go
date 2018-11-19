package middleware

import "github.com/gin-gonic/gin"

type container interface {
	Version() string
}

// YogoMiddleware base type of middleware
// implmeentation in Yoyo
type YogoMiddleware interface {
	Handler() func(c *gin.Context)
}

package server

import (
	"fmt"
	"yogo/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// YogoServerMiddlewares is middleware
// that need to use with server instance
type YogoServerMiddlewares struct {
	SessionTokenValidate middleware.YogoMiddleware
}

type context interface {
	Middlewares() YogoServerMiddlewares
	Version() string
}

// CreateServerInstance factory func
// for create new server
func CreateServerInstance(ctx context) {
	fmt.Println(ctx.Version())

	r := gin.Default()
	r.Use(ctx.Middlewares().SessionTokenValidate.Handler(ctx))

	r.GET("/version", func(c *gin.Context) {
		user := middleware.GetSessionFromRequestContext(c)
		fmt.Println(user)
		c.String(200, user.ID)
	})

	r.Run(":8080")
}
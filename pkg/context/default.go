package context

import (
	"yogo/pkg/middleware"
	"yogo/pkg/server"
)

// YogoContext will composite every
// pkg context to create new one
// that satisfy every package
type YogoContext struct {
	version     string
	middlewares server.YogoServerMiddlewares
}

// Version return string of current API version
func (c *YogoContext) Version() string {
	return c.version
}

// Middlewares every middleware implementation
// to use with server
func (c *YogoContext) Middlewares() server.YogoServerMiddlewares {
	return c.middlewares
}

// CreateContext create new YogoContext
// this context will use as main context
// dependency as production staging
func CreateContext() *YogoContext {

	// Create new YogoContext
	// and encapsulate init data
	return &YogoContext{
		version: "1.0.0",
		middlewares: server.YogoServerMiddlewares{
			SessionTokenValidate: middleware.CreateSessionMiddleware(),
		},
	}

}

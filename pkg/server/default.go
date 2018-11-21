package server

import (
	"net/http"
	"yogo/pkg/middleware"
	"yogo/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// YogoServerMiddlewares is middleware
// that need to use with server instance
type YogoServerMiddlewares struct {
	SessionTokenValidate middleware.YogoMiddleware
}

type container interface {
	Middlewares() YogoServerMiddlewares
	Models() models.Models

	// GraphQLHandler must return GraphQL endpoint handler and playground handler
	GraphQLHandler() (http.HandlerFunc, http.HandlerFunc)
	ViewsHandler() http.HandlerFunc
}

// CreateServerInstance factory
// for create new server ja
func CreateServerInstance(ctx container) *gin.Engine {
	r := gin.Default()
	r.Use(ctx.Middlewares().SessionTokenValidate.Handler())

	r.GET("/version", func(c *gin.Context) {
		userModel := ctx.Models().UserModel
		user := userModel.Create()
		user.Email = "rungsikorn@me.com"

		result, err := user.Save()
		if err != nil {
			panic(err)
		}

		c.String(200, result.(objectid.ObjectID).Hex())
	})

	graphql, playground := ctx.GraphQLHandler()
	r.GET("/graphql", func(c *gin.Context) {
		playground.ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/graphql", func(c *gin.Context) {
		graphql.ServeHTTP(c.Writer, c.Request)
	})

	r.NoRoute(func(c *gin.Context) {
		ctx.ViewsHandler().ServeHTTP(c.Writer, c.Request)
	})

	return r
}

package server

import (
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
}

// CreateServerInstance factory func
// for create new server
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

	return r
}

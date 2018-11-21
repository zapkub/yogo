package di

import (
	"fmt"
	"net/http"
	"yogo/pkg/database"
	"yogo/pkg/graphql"
	"yogo/pkg/middleware"
	"yogo/pkg/models"
	"yogo/pkg/server"
	"yogo/pkg/views"

	"github.com/caarlos0/env"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// YogoContainer will composite every
// pkg context to create new one
// that satisfy every package
type YogoContainer struct {
	version        string
	databaseConfig *database.Config
	viewConfig     *views.Config

	middlewares server.YogoServerMiddlewares

	db     *mongo.Database
	models *models.Models

	graphQLPlaygroundHandler http.HandlerFunc
	graphQLEndpointHandler   http.HandlerFunc
	viewsHandler             http.HandlerFunc
}

// Version return string of current API version
func (c *YogoContainer) Version() string {
	return c.version
}

// Middlewares every middleware implementation
// to use with server
func (c *YogoContainer) Middlewares() server.YogoServerMiddlewares {
	return c.middlewares
}

// DatabaseConfig provide config from .env to
// database factory
func (c *YogoContainer) DatabaseConfig() database.Config {
	if c.databaseConfig == nil {
		panic("DatabaseConfig not found")
	}
	return *c.databaseConfig
}

// DB database connection instance
func (c *YogoContainer) DB() *mongo.Database {
	if c.db == nil {
		fmt.Println("Create DB Connection...")
		mongoDBConnection, err := database.CreateMongoDBClient(c)
		if err != nil {
			fmt.Printf("Create db connection error")
			panic(err)
		}
		c.db = mongoDBConnection
	}
	return c.db
}

// Models singleton models instance
func (c *YogoContainer) Models() models.Models {

	if c.models == nil {
		c.models = models.CreateNewModels(c)
	}

	return *c.models
}

// GraphQLHandler return GraphQL endpoint and playground handler function
func (c *YogoContainer) GraphQLHandler() (http.HandlerFunc, http.HandlerFunc) {

	if c.graphQLEndpointHandler == nil || c.graphQLPlaygroundHandler == nil {
		c.graphQLEndpointHandler = graphql.CreateGraphQLHandler(c)
		c.graphQLPlaygroundHandler = graphql.CreateGraphQLPlaygroundHandler()
	}

	return c.graphQLEndpointHandler, c.graphQLPlaygroundHandler
}

func (c *YogoContainer) ViewConfig() views.Config {
	if c.viewConfig == nil {
		c.viewConfig = &views.Config{}
		err := env.Parse(c.viewConfig)
		if err != nil {
			panic(err)
		}
	}
	return *c.viewConfig
}

func (c *YogoContainer) ViewsHandler() http.HandlerFunc {
	if c.viewsHandler == nil {
		c.viewsHandler = views.CreateViewsHandler(c)

	}
	return c.viewsHandler
}

// CreateDependenciesContainer create new YogoContext
// this context will use as main context
// dependency as production staging
func CreateDependenciesContainer() *YogoContainer {
	var mainContext *YogoContainer

	middlewares := server.YogoServerMiddlewares{
		SessionTokenValidate: middleware.CreateSessionMiddleware(mainContext),
	}

	// create database config
	databaseConfig := database.Config{}
	err := env.Parse(&databaseConfig)
	if err != nil {
		panic("Config parser failed...")
	}

	// Create new YogoContext
	// and encapsulate init data
	mainContext = &YogoContainer{
		version:        "1.0.0",
		middlewares:    middlewares,
		databaseConfig: &databaseConfig,
	}

	mainContext.DB()
	return mainContext
}

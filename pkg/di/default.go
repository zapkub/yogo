package di

import (
	"fmt"
	"yogo/pkg/database"
	"yogo/pkg/middleware"
	"yogo/pkg/models"
	"yogo/pkg/server"

	"github.com/caarlos0/env"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// YogoContext will composite every
// pkg context to create new one
// that satisfy every package
type YogoContext struct {
	version        string
	middlewares    server.YogoServerMiddlewares
	databaseConfig *database.Config
	db             *mongo.Database
	models         *models.Models
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

// DatabaseConfig provide config from .env to
// database factory
func (c *YogoContext) DatabaseConfig() database.Config {
	if c.databaseConfig == nil {
		panic("DatabaseConfig not found")
	}
	return *c.databaseConfig
}

// DB database connection instance
func (c *YogoContext) DB() *mongo.Database {
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
func (c *YogoContext) Models() models.Models {

	if c.models == nil {
		c.models = models.CreateNewModels(c)
	}

	return *c.models
}

// CreateDependenciesContainer create new YogoContext
// this context will use as main context
// dependency as production staging
func CreateDependenciesContainer() *YogoContext {
	var mainContext *YogoContext

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
	mainContext = &YogoContext{
		version:        "1.0.0",
		middlewares:    middlewares,
		databaseConfig: &databaseConfig,
	}

	mainContext.DB()
	return mainContext
}

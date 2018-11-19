package database

// Config database config
type Config struct {
	MongoDBURL string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017"`
}

type container interface {
	DatabaseConfig() Config
}

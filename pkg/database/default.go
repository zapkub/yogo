package database

// Config database config
type Config struct {
	MongoDBURL string `env:"MONGODB_URL"`
}

type context interface {
	DatabaseConfig() Config
}

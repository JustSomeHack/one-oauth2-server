package mongoconfig

import "os"

// MongoConfig struct for storing configuration values
type MongoConfig struct {
	MongoURL string
	DBName   string
	DBUser   string
	DBPass   string
}

// LoadMongoConfig pulls configuration from env
func LoadMongoConfig() *MongoConfig {

	return &MongoConfig{
		MongoURL: os.Getenv("MONGO_URL"),
		DBName:   os.Getenv("DB_NAME"),
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASS"),
	}
}

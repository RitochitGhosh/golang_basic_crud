package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("Missing required environment variable")
	}

	return val, nil
}

func Load() (Config, error) {

	// godotenv.Load() -> reads .env and sets them into the process env
	// os.Getenv() -> reads these values
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env")
		return Config{}, fmt.Errorf("Failed to load .env: %w", err)
	}

	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil
}

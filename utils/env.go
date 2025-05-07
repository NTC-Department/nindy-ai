package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	checkEnv()
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func checkEnv() {
	requiredEnvVars := []string{
		"BOT_TOKEN",
		"GEMINI_API_KEY",
		"CHATBOT_CHANNEL_ID",
	}

	for _, envVar := range requiredEnvVars {
		if GetEnv(envVar) == "" {
			panic(envVar + " is not set in the environment")
		}
	}
}

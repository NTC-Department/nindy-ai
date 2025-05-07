package utils

import (
	"log"
	"os"
	"strings"

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

func GetEnvWithMultipleValue(key string) []string {
	value := os.Getenv(key)
	lines := strings.Split(strings.TrimSpace(value), "\n")
	var result []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			result = append(result, trimmedLine)
		}
	}
	return result
}

func checkEnv() {
	requiredEnvVars := []string{
		"BOT_TOKEN",
		"GEMINI_API_KEY",
		"CHATBOT_CHANNEL_ID",
		"ADMIN_USER_ID",
		"API_KEY_ROTATION_FREQUENCY",
	}

	for _, envVar := range requiredEnvVars {
		if GetEnv(envVar) == "" {
			panic(envVar + " is not set in the environment")
		}
	}
}

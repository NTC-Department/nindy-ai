package external

import (
	"context"
	"log"
	"nindychat/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var rotationFrequency = 20
var apiKeys []string
var apiKeyIndex = 0
var usageCounter = 0

var currentClient *genai.Client
var currentModel *genai.GenerativeModel

func InitializeGeminiEnv() {
	apiKeys = utils.GetEnvWithMultipleValue("GEMINI_API_KEY")
	if len(apiKeys) == 0 {
		panic("No Gemini API keys found in environment")
	}
	log.Printf("Found %d Gemini API keys\n", len(apiKeys))

	frequency, err := strconv.Atoi(utils.GetEnv("API_KEY_ROTATION_FREQUENCY"))
	if err != nil {
		log.Fatalf("Invalid API_KEY_ROTATION_FREQUENCY value: %v", err)
	} else {
		rotationFrequency = frequency
	}
}

func InitializeGemini() {
	ctx := context.Background()
	apiKey := getNextAPIKeyInternal()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error initializing Gemini client with initial key: %v", err)
		return
	}
	currentClient = client
	currentModel = client.GenerativeModel("gemini-2.0-flash")
	currentModel.SetTemperature(0.7)
	currentModel.SetMaxOutputTokens(1000)

	log.Println("Gemini model initialized")
	go waitForShutdown()
}

func getNextAPIKeyInternal() string {
	key := apiKeys[apiKeyIndex%len(apiKeys)]
	apiKeyIndex++
	return key
}

func RotateGeminiAPIKey() {
	if len(apiKeys) <= 1 {
		return
	}

	log.Println("Rotating Gemini API Key...")
	if currentClient != nil {
		currentClient.Close()
	}

	InitializeGemini()
}

func GetGeminiModel() *genai.GenerativeModel {
	if currentModel == nil {
		panic("Gemini model failed to initialize")
	}

	return currentModel
}

func IncrementGeminiUsage() {
	usageCounter++
	if usageCounter >= rotationFrequency {
		log.Println("Max usage reached, rotating API key...")
		usageCounter = 0
		RotateGeminiAPIKey()
	}
}

func CloseGeminiClient() {
	if currentClient != nil {
		log.Println("Closing Gemini client...")
		currentClient.Close()
		currentClient = nil
	}
}

func waitForShutdown() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc
	CloseGeminiClient()
	os.Exit(0)
}

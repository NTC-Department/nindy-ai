package external

import (
	"context"
	"log"
	"nindychat/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var geminiModel *genai.GenerativeModel
var geminiClient *genai.Client

func InitializeGemini() {
	ctx := context.Background()
	apiKey := utils.GetEnv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	geminiClient = client

	model := client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.7)
	model.SetMaxOutputTokens(1000)

	geminiModel = model

	log.Println("Gemini model initialized successfully")
	go waitForShutdown()
}

func GetGeminiModel() *genai.GenerativeModel {
	if geminiModel == nil {
		panic("Gemini model failed to initialize")
	}

	return geminiModel
}

func CloseGeminiClient() {
	if geminiClient != nil {
		log.Println("Closing Gemini client...")
		geminiClient.Close()
	}
}

func waitForShutdown() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc
	CloseGeminiClient()
	os.Exit(0)
}

package config

import (
	"log"
	"nindychat/feature/chatbot"

	"github.com/bwmarrin/discordgo"
)

func InitializeHandler(dg *discordgo.Session) {
	log.Println("Adding Handler...")
	dg.AddHandler(chatbot.Handler)
	log.Println("✓ Handler added")
}

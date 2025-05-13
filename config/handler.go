package config

import (
	"log"
	"nindychat/feature/admin"
	"nindychat/feature/chatbot"

	"github.com/bwmarrin/discordgo"
)

func InitializeHandler(dg *discordgo.Session) {
	log.Println("Adding Handler...")
	dg.AddHandler(chatHandler)
	log.Println("✓ Handler added")
}

func chatHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if admin.Handler(s, m) {
		return
	}

	chatbot.Handler(s, m)
}

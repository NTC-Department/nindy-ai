package config

import (
	"nindychat/feature/chatbot"

	"github.com/bwmarrin/discordgo"
)

func InitializeHandler(dg *discordgo.Session) {
	dg.AddHandler(chatbot.Handler)
}

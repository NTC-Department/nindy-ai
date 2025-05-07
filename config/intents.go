package config

import "github.com/bwmarrin/discordgo"

func InitializeIntents(dg *discordgo.Session) {
	var intents = []discordgo.Intent{
		discordgo.IntentsGuildMessages,
		discordgo.IntentsGuildEmojis,
	}

	for _, intent := range intents {
		dg.Identify.Intents |= intent
	}
}

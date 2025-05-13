package admin

import (
	"nindychat/external"
	"nindychat/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	prefix := utils.GetEnv("BOT_PREFIX")

	if m.Author.ID == utils.GetEnv("ADMIN_USER_ID") {
		if strings.HasPrefix(m.Content, prefix+"rotate") {
			external.RotateGeminiAPIKey()
			s.ChannelMessageSend(m.ChannelID, "Rotated Gemini API Key!")
			return true
		}
	}

	return false
}

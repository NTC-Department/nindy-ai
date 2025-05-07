package chatbot

import (
	"nindychat/external"
	"nindychat/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// ADMIN OVERRIDE COMMAND
	if m.Author.ID == utils.GetEnv("ADMIN_USER_ID") {
		if strings.HasPrefix(m.Content, "-rotate") {
			external.RotateGeminiAPIKey()
			s.ChannelMessageSend(m.ChannelID, "Rotated Gemini API Key!")
			return
		}
	}

	// for bot usage outside of the specific channel
	if strings.HasPrefix(m.Content, "<@"+s.State.User.ID) {
		NewChat(s, m).Chat()
		return
	}

	// for bot usage inside the specific channel
	if m.ChannelID == utils.GetEnv("CHATBOT_CHANNEL_ID") {
		if strings.HasPrefix(m.Content, "// ") {
			return
		}
		NewChat(s, m).Chat()
		return
	}
}

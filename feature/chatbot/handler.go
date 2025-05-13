package chatbot

import (
	"nindychat/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if strings.HasPrefix(m.Content, "<@"+s.State.User.ID) {
		NewChat(s, m).Chat()
		return true
	}

	if m.ChannelID == utils.GetEnv("CHATBOT_CHANNEL_ID") {
		if strings.HasPrefix(m.Content, "// ") {
			return false
		}
		NewChat(s, m).Chat()
		return true
	}

	return false
}

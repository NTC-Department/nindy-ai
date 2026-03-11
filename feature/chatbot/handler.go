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

	allowedChannels := strings.Split(utils.GetEnv("CHATBOT_CHANNEL_ID"), ",")
	isAllowedChannel := false
	for _, id := range allowedChannels {
		if strings.TrimSpace(id) == m.ChannelID {
			isAllowedChannel = true
			break
		}
	}

	if isAllowedChannel {
		if strings.HasPrefix(m.Content, "// ") {
			return false
		}
		NewChat(s, m).Chat()
		return true
	}

	return false
}

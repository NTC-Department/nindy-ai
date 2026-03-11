package chatbot

import (
	"nindychat/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	isMentioned := strings.HasPrefix(m.Content, "<@"+s.State.User.ID)
	isReplyToBot := m.ReferencedMessage != nil && m.ReferencedMessage.Author != nil && m.ReferencedMessage.Author.ID == s.State.User.ID

	if isMentioned || isReplyToBot {
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

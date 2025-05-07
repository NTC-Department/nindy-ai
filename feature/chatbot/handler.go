package chatbot

import (
	"nindychat/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "<@"+s.State.User.ID) {
		NewChat(s, m).Chat()
	} else if m.ChannelID == utils.GetEnv("CHATBOT_CHANNEL_ID") {
		NewChat(s, m).Chat()
	}
}

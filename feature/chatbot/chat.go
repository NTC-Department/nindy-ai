package chatbot

import (
	"context"
	"log"
	"nindychat/external"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
)

type chat struct {
	session *discordgo.Session
	msg     *discordgo.MessageCreate
}

func NewChat(s *discordgo.Session, m *discordgo.MessageCreate) *chat {
	return &chat{
		session: s,
		msg:     m,
	}
}

func (c *chat) Chat() {
	geminiModel := external.GetGeminiModel()
	ctx := context.Background()

	userName := c.msg.Member.Nick
	if userName == "" {
		userName = c.msg.Author.Username
	}

	err := c.session.MessageReactionAdd(c.msg.ChannelID, c.msg.ID, "🔄")
	if err != nil {
		log.Printf("Failed to add reaction: %v", err)
	}

	question := strings.TrimPrefix(c.msg.Content, "<@"+c.session.State.User.ID+">")
	question = strings.TrimSpace(question)

	chatHistory := getChatHistory(c.msg.ChannelID)
	appendChatHistory(c.msg.ChannelID, userName, question)

	question = BuildPrompt(chatHistory, userName, question)

	resp, err := geminiModel.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
		c.session.ChannelMessageSendReply(
			c.msg.ChannelID,
			"[!] Error: "+err.Error(),
			&discordgo.MessageReference{
				MessageID: c.msg.ID,
				ChannelID: c.msg.ChannelID,
				GuildID:   c.msg.GuildID,
			},
		)
		c.session.MessageReactionRemove(c.msg.ChannelID, c.msg.ID, "🔄", "@me")
		external.RotateGeminiAPIKey()
		return
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		switch p := part.(type) {
		case genai.Text:
			text := string(p)
			appendChatHistory(c.msg.ChannelID, "Nindy Luzie", text)

			c.session.ChannelMessageSendReply(
				c.msg.ChannelID,
				text,
				&discordgo.MessageReference{
					MessageID: c.msg.ID,
					ChannelID: c.msg.ChannelID,
					GuildID:   c.msg.GuildID,
				},
			)
		default:
			log.Printf("Unknown part type: %T\n", p)
			c.session.ChannelMessageSendReply(
				c.msg.ChannelID,
				"[!] Non-text part received",
				&discordgo.MessageReference{
					MessageID: c.msg.ID,
					ChannelID: c.msg.ChannelID,
					GuildID:   c.msg.GuildID,
				},
			)
		}
	}

	external.IncrementGeminiUsage()

	err = c.session.MessageReactionRemove(c.msg.ChannelID, c.msg.ID, "🔄", "@me")
	if err != nil {
		log.Printf("Failed to remove reaction: %v", err)
	}
}

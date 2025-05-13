package chatbot

import (
	"fmt"
	"nindychat/database/redis"
	"time"
)

const CHAT_HISTORY_STORE_TIME = time.Minute * 10

func getChatHistory(channelID string) string {
	redisKey := "chat_history:" + channelID
	history, err := redis.Get(redisKey)
	if err != nil {
		fmt.Printf("Error getting chat history for %s: %v\n", channelID, err)
		return ""
	}
	return history
}

func appendChatHistory(channelID, userNickname, chat string) {
	redisKey := "chat_history:" + channelID
	newMessage := fmt.Sprintf("%s:%s", userNickname, chat)

	err := redis.Append(redisKey, newMessage+"\n")
	if err != nil {
		fmt.Printf("Error appending to chat history for %s: %v\n", channelID, err)
		redis.Set(redisKey, newMessage+"\n", CHAT_HISTORY_STORE_TIME)
	} else {
		redis.Expire(redisKey, CHAT_HISTORY_STORE_TIME)
	}
}

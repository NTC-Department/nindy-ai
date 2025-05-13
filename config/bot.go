package config

import (
	"log"
	"nindychat/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func InitializeBot() {
	log.Println("===== Initialize Discord Bot =====")
	dg, err := discordgo.New("Bot " + utils.GetEnv("BOT_TOKEN"))
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	InitializeHandler(dg)
	InitializeIntents(dg)

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	log.Println(dg.State.User.Username + " is now online!")
	waitForShutdown(dg)
}

func waitForShutdown(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}

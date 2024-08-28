package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Print("Hello!")

	sess, err := discordgo.New("Bot MTI3ODQ4NjczOTAxNDUyMDkwMw.G4-3H5.HmYjnt2XcS6MsWVG0ZCv8NAtJmrfDd0Q7kpdsU")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.ToLower(m.Content) == "neublox" {
			s.ChannelMessageSend(m.ChannelID, "Build powerful machine learning models using plain language. Our software enables anyone to create and deploy AI effortlessly. Revolutionize your approach to machine learning — no coding required.")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
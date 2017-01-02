package main

import (
	"flag"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	botToken := flag.String("token", "", "bot token to use telegram bot api")
	// dataDir := flag.String("datadir", "/var/data", "data dir to save files")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	if *botToken == "" {
		log.Println("botToken must be set")
		os.Exit(0)
	}

	bot, errnba := tgbotapi.NewBotAPI(*botToken)
	if errnba != nil {
		log.Fatalf("%s", errnba)
	}
	bot.Debug = *debug

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updatesChan, erruc := bot.GetUpdatesChan(u)
	if erruc != nil {
		log.Fatal(erruc)
	}

	for {
		select {
		case update := <-updatesChan:
			if update.Message == nil {
				continue
			}
			strs := strings.Split(update.Message.Text, " ")
			command := strings.Split(strings.ToLower(strs[0]), "@")[0]
			if command == "" {
				continue
			}
			if command == "/help" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "HELP STUB")
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				bot.Send(msg)
			}
		}
	}
}

package main

import (
	"flag"
	tgClient "go-read-adviser-bot/clients/telegram"
	event_consumer "go-read-adviser-bot/consumer/event-consumer"
	"go-read-adviser-bot/events/telegram"
	"go-read-adviser-bot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	bathSize    = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("")
	}

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}

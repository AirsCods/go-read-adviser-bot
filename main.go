package main

import (
	"flag"
	"go-read-adviser-bot/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient = telegram.New(tgBotHost, mustToken())

	// Отправляет запросы и получает новые события
	// fetcher = fetcher.New()

	// После обработки отправляет нвоые сообщения
	// processor = processor.New()

	//consumer.Start(fetcher, processor)
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

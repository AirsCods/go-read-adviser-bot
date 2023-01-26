package main

import (
	"flag"
	"log"
)

func main() {
	// Получение токена из флага
	token = mustToken()

	// tgClient = telegram.New(token)

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

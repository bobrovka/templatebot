package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bobrovka/templatebot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	port := "8090"

	cfg := config.Config{
		TgToken: "your token",
		Webhook: "ngrok https url",
	}

	go func() {
		_ = http.ListenAndServe(":"+port, nil)
	}()

	bot, err := tgbotapi.NewBotAPI(cfg.TgToken)
	if err != nil {
		log.Fatal("cannot new bot ", err)
	}
	log.Println("Bot created")

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.Webhook))
	if err != nil {
		log.Fatal("cannot new bot ", err)
	}
	log.Println("Webhook created")

	for update := range bot.ListenForWebhook("/") {
		fmt.Println("Got message: ", update.Message.Text)
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "A: "+update.Message.Text)); err != nil {
			log.Println(err)
		}
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bobrovka/templatebot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	port := os.Getenv("PORT")

	cfg := config.Config{
		TgToken: "1375298760:AAGUhlCpoiEp5PXcN_b38nrWN40-r0jDKQ0",
		Webhook: "https://somedifferenttestbot.herokuapp.com",
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
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "A: "+update.Message.Text)); err != nil {
			log.Println(err)
		}
	}
	_ = port
}

package main

import (
	"log"
	"os"

	bot "blgn.dev/kaihen/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.BotToken = os.Getenv("DISCORD_BOT_TOKEN")
	bot.Run()
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
	App   string
	Guild string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	app := os.Getenv("DISCORD_APPLICATION_ID")
	guild := os.Getenv("DISCORD_GUILD_ID")

	return Config{
		Token: token,
		App:   app,
		Guild: guild,
	}
}

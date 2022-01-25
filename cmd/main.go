package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
	"yosukebot/internal/handler"

	tb "gopkg.in/tucnak/telebot.v3"
)

func main() {
	//log.SetFlags(log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	log.Println("запущено работает")

	if err != nil {
		log.Fatal(err)
		return
	}

	h, err := handler.NewHandler(b)
	if err != nil {
		log.Fatal(err)
		return
	}

	h.Start()
}

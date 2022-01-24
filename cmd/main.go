package main

import (
	"log"
	"time"
	"yosukebot/internal/handler"

	tb "gopkg.in/tucnak/telebot.v3"
)

func main() {
	//log.SetFlags(log.Lshortfile)

	b, err := tb.NewBot(tb.Settings{
		Token:  "5010011691:AAE_2wG6GZq9bchf27LKHZxxFo3oxVkjn44",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

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

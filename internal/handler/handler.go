package handler

import (
	tb "gopkg.in/tucnak/telebot.v3"
)

type Handler struct {
	Bot      *tb.Bot
	Commands Commands
}

func NewHandler(b *tb.Bot) (*Handler, error) {
	commands := Commands{}

	passwordCommand := Command{
		Title:       "Password generation",
		Description: "p [1-100] [login]",
		Aliases:     []string{"p", "pass", "password"},
	}

	shorterCommand := Command{
		Title:       "Short URL generator",
		Description: "s [URL]",
		Aliases:     []string{"s", "short"},
	}

	commands.Add(passwordCommand)
	commands.Add(shorterCommand)

	h := &Handler{
		Bot:      b,
		Commands: commands,
	}
	return h, nil
}

func (h *Handler) Start() {
	//h.Bot.Handle("/music", h.Music)

	h.Bot.Handle("/start", h.Password)

	h.Bot.Handle(tb.OnQuery, h.Query)

	h.Bot.Start()
}

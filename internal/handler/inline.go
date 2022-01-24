package handler

import (
	tb "gopkg.in/tucnak/telebot.v3"
	"strings"
)

func (h Handler) Query(c tb.Context) error {
	command := strings.Split(c.Data(), " ")[0]
	args := c.Args()

	switch command {
	case "p", "pass", "password":
		return h.OnPassword(c, args[1:])
	case "s", "short":
		return h.OnShorter(c, args[1:])
	case "pm":
		return h.OnPM(c, args[1:])
	default:
		return h.OnHelp(c)
	}
}

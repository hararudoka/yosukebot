package handler

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v3"
)

func (h Handler) OnPM(c tb.Context, args []string) error {
	results := make(tb.Results, 1)

	result := &tb.ArticleResult{}

	username := args[0]

	text := fmt.Sprint(args[1:])

	{
		result = &tb.ArticleResult{
			Title:       "Your PM: ",
			Text:        text,
			Description: username,
		}
	}

	results[0] = result

	err := c.Answer(&tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
	return err
}

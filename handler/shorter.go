package handler

import (
	tb "gopkg.in/tucnak/telebot.v3"
)

func (h Handler) OnShorter(c tb.Context, args []string) error {
	results := make(tb.Results, 1)

	result := &tb.ArticleResult{}

	url := args[0]

	shortLink := short(url)

	{
		result = &tb.ArticleResult{
			Title:       "Your short link: ",
			Text:        shortLink,
			Description: shortLink,
		}
	}

	results[0] = result

	err := c.Answer(&tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
	return err
}

func short(u string) string {
	//TODO make API of shorter
	return ""
}


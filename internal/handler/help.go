package handler

import (
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v3"
)

func (h *Handler) OnHelp(c tb.Context) error {
	results := tb.Results{}

	for _, e := range h.Commands {
		result := &tb.ArticleResult{
			Title:       e.Title,
			Description: e.Description,
			Text:        e.Description,
		}
		result.SetResultID(uuid.New().String())
		results = append(results, result)
	}

	huet := tb.QueryResponse{
		Results:   results,
		CacheTime: -1,
	}

	err := c.Answer(&huet)

	return err
}

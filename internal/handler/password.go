package handler

import (
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v3"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (h Handler) Password(m tb.Context) error {
	args := strings.Split(m.Text(), " ")

	message := ""
	if len(args) == 1 {
		message = passGen(12)
	} else if len(args) == 2 {
		n, ok := validate(args[1])
		if ok {
			message = passGen(n)
		}
		message = "ошибка"
	}

	err := m.Send(message)
	return err
}

func passGen(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func validate(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	if n > 100 || n < 1 {
		return 0, false
	}
	return n, true
}

func (h Handler) OnPassword(c tb.Context, args []string) error {
	results := tb.Results{}

	bad := &tb.ArticleResult{
		Title:       "Error: ",
		Text:        "Wrong number",
		Description: "Wrong number",
	}

	if len(args) == 0 {
		for i := 0; i < 5; i++ {
			rand.Seed(time.Now().Unix() + int64(i))
			p := passGen(12)
			result := &tb.ArticleResult{
				Title:       "Your password: ",
				Text:        p,
				Description: p,
			}
			result.SetResultID(uuid.New().String())
			results = append(results, result)
		}
	} else if len(args) == 1 {
		if n, ok := validate(args[0]); ok {
			for i := 0; i < 5; i++ {
				rand.Seed(time.Now().Unix() + int64(i))
				p := passGen(n)
				result := &tb.ArticleResult{
					Title:       "Your password: ",
					Text:        p,
					Description: p,
				}
				result.SetResultID(uuid.New().String())
				results = append(results, result)
			}
		} else {
			results = append(results, bad)
		}
	} else if len(args) == 2 {
		if n, ok := validate(args[0]); ok {
			for i := 0; i < 5; i++ {
				rand.Seed(time.Now().Unix() + int64(i))
				p := passGen(n)
				result := &tb.ArticleResult{
					Title:       "Your password: ",
					Text:        args[1] + ":" + p,
					Description: args[1] + ":" + p,
				}
				result.SetResultID(uuid.New().String())
				results = append(results, result)
			}
		} else {
			results = append(results, bad)
		}
	} else {
		results = append(results, bad)
	}

	err := c.Answer(&tb.QueryResponse{
		Results:   results,
		CacheTime: 2,
	})

	return err
}

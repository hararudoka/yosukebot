package handler

import (
	"bytes"
	"encoding/json"
	tb "gopkg.in/tucnak/telebot.v3"
	"io/ioutil"
	"net/http"
)

func (h Handler) OnShorter(c tb.Context, args []string) error {
	results := make(tb.Results, 1)

	result := &tb.ArticleResult{}

	url := args[0]

	shortLink, err := short(url)
	if err != nil {
		return err
	}

	{
		result = &tb.ArticleResult{
			Title:       "Your short link: ",
			Text:        shortLink,
			Description: shortLink,
		}
	}

	results[0] = result

	err = c.Answer(&tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
	return err
}

func short(u string) (string,error) {
	postBody, _ := json.Marshal(map[string]string{
		"url":  u,
	})
	responseBody := bytes.NewBuffer(postBody)
	r, err := http.Post("https://h.mxf.su/api", "application/json", responseBody)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return " ", err
	}
	return string(body[1:len(body)-2]), nil
}


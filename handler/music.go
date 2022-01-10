package handler

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os/exec"
	"strings"
)

func (h Handler) Music(um *tb.Message) {
	args := strings.Split(um.Text, " ")

	a := audioFromYM(args[1])
	_, err := h.Bot.Send(um.Sender, a)
	if err != nil {
		log.Println(err)
	}
}

func audioFromYM(id string) *tb.Audio {
	cmd := exec.Command("youtube-dl", "-x", id, "--audio-format", "mp3", "-o", `%(id)s.%(ext)s'`)

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}

	cmd2 := exec.Command("youtube-dl --get-filename -o '%(title)s' BaW_jenozKc") // если байты youtube-dl "abc" -o
	err = cmd2.Start()
	if err != nil {
		panic(err)
	}

	err = cmd2.Wait()
	if err != nil {
		panic(err)
	}

	var name []byte
	cmd2.Stdout.Write(name)

	fmt.Println()
	defer exec.Command("rm", id+".mp3")

	gopherMusic := &tb.Audio{File: tb.FromDisk(id + ".mp3"), Title: "ыыы"}

	return gopherMusic
}

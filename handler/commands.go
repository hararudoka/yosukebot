package handler

import "strings"

type Command struct {
	Title       string
	Description string
	Aliases     []string
	Inline      bool
}

type Commands []Command

func (cs *Commands) Add(c Command) {
	*cs = append(*cs, c)
}

func (cs Commands) Get(s string) Command {
	for _, e := range cs{
		if ok, n := e.IsIn(s); ok {
			return cs[n]
		}
	}
	return Command{}
}

func (c Command) IsIn(s string) (bool, int) {
	com := strings.Split(s, " ")
	for i, e := range c.Aliases {
		if e == com[0] {
			return true, i
		}
	}
	return false, 0
}


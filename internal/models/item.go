package models

import "time"

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DoneAt      string `json:"done_at"`
	CreatedAt   string `json:"created_at"`
}

func (i *Item) MarkAsDone() {
	i.DoneAt = time.Now().Format(time.RFC822)
}

func (i *Item) MarkAsUndone() {
	i.DoneAt = ""
}

func NewItem() *Item {
	return &Item{
		CreatedAt: time.Now().Format(time.RFC822),
	}
}

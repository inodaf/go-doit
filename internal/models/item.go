package models

import (
	"errors"
	"time"
)

var ErrMissingTitleAdd = errors.New("Models/Item: Missing 'title'")

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

func NewItem(title string) (*Item, error) {
	if len(title) == 0 {
		return nil, ErrMissingTitleAdd
	}

	return &Item{
		Title:     title,
		CreatedAt: time.Now().Format(time.RFC822),
	}, nil
}

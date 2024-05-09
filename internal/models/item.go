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
	UpdatedAt   string `json:"update_at"`
}

func (i *Item) MarkAsDone() {
	i.UpdatedAt = time.Now().Format(time.RFC822)
	i.DoneAt = time.Now().Format(time.RFC822)
}

func (i *Item) MarkAsUndone() {
	i.DoneAt = ""
	i.UpdatedAt = time.Now().Format(time.RFC822)
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

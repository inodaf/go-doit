package models

import (
	"errors"
	"time"
)

var ErrMissingTitleAdd = errors.New("models/item: missing 'title'")

type Item struct {
	Id          int `json:"id" sql:"id"`
	Title       string `json:"title" sql:"title"`
	Description string `json:"description" sql:"description"`
	DoneAt      string `json:"done_at" sql:"done_at"`
	CreatedAt   string `json:"created_at" sql:"created_at"`
	UpdatedAt   string `json:"update_at" sql:"update_at"`
}

func (i *Item) MarkAsDone() {
	i.UpdatedAt = time.Now().Format(time.DateTime)
	i.DoneAt = time.Now().Format(time.DateTime)
}

func (i *Item) MarkAsUndone() {
	i.DoneAt = ""
	i.UpdatedAt = time.Now().Format(time.DateTime)
}

func NewItem(title string) (*Item, error) {
	if len(title) == 0 {
		return nil, ErrMissingTitleAdd
	}

	return &Item{
		Title:     title,
		CreatedAt: time.Now().Format(time.DateTime),
	}, nil
}

package models

import (
	"errors"
	"time"
)

var ErrMissingTitleAdd = errors.New("models/item: missing 'title'")

type Item struct {
	Id          int       `json:"id" sql:"id"`
	Title       string    `json:"title" sql:"title"`
	Description string    `json:"description" sql:"description"`
	DoneAt      time.Time `json:"done_at" sql:"done_at"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedAt   time.Time `json:"update_at" sql:"update_at"`
}

func (i *Item) MarkAsDone() {
	i.UpdatedAt = time.Now()
	i.DoneAt = time.Now()
}

func (i *Item) MarkAsUndone() {
	i.DoneAt = time.Time{}
	i.UpdatedAt = time.Now()
}

func NewItem(title string) (*Item, error) {
	if len(title) == 0 {
		return nil, ErrMissingTitleAdd
	}

	return &Item{
		Title:     title,
		CreatedAt: time.Now(),
	}, nil
}

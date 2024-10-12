package models

import (
	"errors"
	"time"
)

var ErrMissingTitleAdd = errors.New("models/item: missing 'title'")

type Item struct {
	Id          int
	Title       string
	Description string
	DoneAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

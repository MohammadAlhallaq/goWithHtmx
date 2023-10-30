package model

import (
	"errors"
)

var (
	ErrItemNotFound    = errors.New("the item was not found in the repository")
	ErrFailedToAddItem = errors.New("failed to add the item to the repository")
	ErrUpdateItem      = errors.New("failed to update the item in the repository")
)

type ItemRepo interface {
	GetAll() ([]Item, error)
	Get(id uint64) (Item, error)
	Add(customer Item) error
	Delete(id uint64) error
	UpdateStatus(id uint64) error
}

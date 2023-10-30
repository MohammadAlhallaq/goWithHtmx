package model

import (
	"errors"
)

var (
	ErrItemNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddItem = errors.New("failed to add the customer to the repository")
	ErrUpdateCItem     = errors.New("failed to update the customer in the repository")
)

type ItemRepo interface {
	GetAll() ([]Item, error)
	Get(id uint64) (Item, error)
	Add(customer Item) error
	Delete(id uint64) error
	UpdateStatus(id uint64) error
}

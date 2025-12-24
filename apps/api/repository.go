package main

import (
	"errors"
)

var ErrContactNotFound = errors.New("contact not found")

type ContactRepository interface {
	Create(contact *contact) error
	List() ([]contact, error)
	Get(id string) (*contact, error)
	Update(id string, payload *contact) error
	Delete(id string) error
}

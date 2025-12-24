package main

import "sync"

type fakeContactRepository struct {
	mu       sync.RWMutex
	contacts map[string]contact
}

func NewFakeContactRepository(seed map[string]contact) ContactRepository {
	return &fakeContactRepository{
		contacts: seed,
	}
}

func (r *fakeContactRepository) Create(payload *contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contacts[payload.ID] = *payload
	return nil
}

func (r *fakeContactRepository) List() ([]contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	contacts := make([]contact, 0, len(r.contacts))
	for _, co := range r.contacts {
		contacts = append(contacts, co)
	}
	return contacts, nil
}

func (r *fakeContactRepository) Get(id string) (*contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	co, ok := r.contacts[id]
	if !ok {
		return nil, ErrContactNotFound
	}
	result := co
	return &result, nil
}

func (r *fakeContactRepository) Update(id string, payload *contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.contacts[id]
	if !ok {
		return ErrContactNotFound
	}

	r.contacts[id] = contact{
		ID:          id,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
	}
	return nil
}

func (r *fakeContactRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.contacts[id]; !ok {
		return ErrContactNotFound
	}

	delete(r.contacts, id)
	return nil
}

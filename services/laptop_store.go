package services

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/jinzhu/copier"
)

// ErrAlreadyExists is returned when a record with the same ID already exist in the store
var ErrAlreadyExists = errors.New("record already exist")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *uploadpb.Laptop) error
	// Find find the laptop by its id
	Find(id string) (*uploadpb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*uploadpb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*uploadpb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *uploadpb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &uploadpb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	// save to store
	store.data[other.Id] = other
	return nil
}

// Find the related store by its id
func (store *InMemoryLaptopStore) Find(id string) (*uploadpb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := &uploadpb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}

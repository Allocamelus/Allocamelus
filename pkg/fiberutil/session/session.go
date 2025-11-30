//go:generate msgp
package session

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrNilKey returned when key has 0 or less length
	ErrNilKey = errors.New("Session: Error 0 length key")
	// ErrNilValue returned when value is nil
	ErrNilValue = errors.New("Session: Error nil value")
	// ErrValueType returned when value type is invalid
	ErrValueType = errors.New("Session: Error invalid value type")
)

// Status of Session
type Status int8

const (
	// Created Session Status
	Created Status = iota
	// NotUpdated Session Status
	NotUpdated
	// Updated Session Status
	Updated
	// Deleted Session Status
	Deleted
)

// Session is session
type Session struct {
	// Created time.Time
	Created time.Time `msg:"created"`
	// Expires time.Time
	Expires time.Time `msg:"expires"`
	// Status of session
	Status Status `msg:"status"`
	// Key session identifier
	Key string `msg:"key"`
	// Data of session
	Data map[string]any `msg:"data"`
	// mu RWMutex
	mu sync.RWMutex
	// store pointer to store
	store *Store
}

// Init add Store and mutex to session
func (s *Session) Init(store *Store) {
	s.mu.Lock()
	s.store = store
	s.mu.Unlock()
}

// Create new Session
func (s *Session) Create(key string) {
	s.mu.Lock()
	s.Key = key
	s.Created = time.Now()
	s.Status = Created
	s.Data = make(map[string]any)
	s.mu.Unlock()
	s.Updated()
}

// Regenerate Session
func (s *Session) Regenerate() {
	s.store.Data.Delete(s.Key)
	s.Create(s.store.Key.Generator(s.store.Key.Length))
}

// Set session value
func (s *Session) Set(key string, value any) error {
	if key == "" {
		return ErrNilKey
	}
	if value == nil {
		return ErrNilValue
	}
	s.mu.Lock()
	s.Data[key] = value
	s.mu.Unlock()
	s.Updated()
	return nil
}

// Updated updates expires and status
func (s *Session) Updated() {
	s.mu.Lock()
	s.Expires = time.Now().Add(s.store.Expiration)
	if s.Status != Created && s.Status != Deleted {
		s.Status = Updated
	}
	s.mu.Unlock()
}

// Get session value
func (s *Session) Get(key string) (any, error) {
	if key == "" {
		return nil, ErrNilKey
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Does Data have any
	if len(s.Data) > 0 {
		if value, ok := s.Data[key]; ok {
			return value, nil
		}
	}
	return nil, ErrNilValue
}

// GetBytes from session
func (s *Session) GetBytes(key string) ([]byte, error) {
	bInter, err := s.Get(key)
	if err != nil {
		return nil, err
	}
	if b, ok := bInter.([]byte); ok {
		return b, nil
	}
	return nil, ErrValueType
}

// Delete session does not regenerate
func (s *Session) Delete() {
	s.store.Data.Delete(s.Key)
	s.mu.Lock()
	s.Status = Deleted
	s.mu.Unlock()
}

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
	// Created Session Status for created
	Created Status = iota
	// NotUpdated Session Status for no change
	NotUpdated
	// Updated Session Status for change
	Updated
	// Deleted Session Status for deletion
	Deleted
)

//go:generate msgp

// Session is session
type Session struct {
	Created time.Time              `msg:"created"`
	Expires time.Time              `msg:"expires"`
	Status  Status                 `msg:"status"`
	Key     string                 `msg:"key"`
	Data    map[string]interface{} `msg:"data"`
	mu      sync.RWMutex
	store   *Store
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
	s.Data = make(map[string]interface{})
	s.mu.Unlock()
	s.Updated()
}

// Regenerate Session
func (s *Session) Regenerate() {
	s.store.Data.Delete(s.Key)
	s.Create(s.store.Key.Generator(s.store.Key.Length))
}

// Set session value
func (s *Session) Set(key string, value interface{}) error {
	if len(key) > 0 {
		if value != nil {
			s.mu.Lock()
			s.Data[key] = value
			s.mu.Unlock()
			s.Updated()
			return nil
		}
		return ErrNilValue
	}
	return ErrNilKey
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
func (s *Session) Get(key string) (interface{}, error) {
	if len(key) > 0 {
		if len(s.Data) > 0 {
			if value, ok := s.Data[key]; ok {
				return value, nil
			}
		}
		return nil, ErrNilValue
	}
	return nil, ErrNilKey
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

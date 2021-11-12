package store

import (
	"KVdb/entity"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type Store struct {
	m map[entity.Key]entity.Value
	sync.RWMutex
}

func ValidTime(t time.Time) bool {
	return t.After(time.Now())
}

// removes a key, with thread safety
func (s *Store) safeRemove(k entity.Key) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, k)
}

func (s *Store) safeRead(k entity.Key) (entity.Value, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.m[k]
	return val, ok
}

// Get checks if a value exists and if it is valid returns it, otherwise removes the value from map
func (s *Store) Get(_ context.Context, k string) (interface{}, error) {
	key := entity.Key(k)

	val, ok := s.safeRead(key)
	if !ok {
		return nil, entity.ErrNotFound
	}
	value, err := val.GetValue()

	// if value is explored remove it
	if err == entity.ErrExpired {
		s.safeRemove(key)
		return nil, entity.ErrExpired
	} else if err != nil {
		return nil, err
	}
	return value, nil
}

// Set sets a set of Key Value to the hash map
func (s *Store) Set(_ context.Context, k string, v interface{}, exp time.Time) error {
	if !ValidTime(exp) {
		return entity.ErrExpired
	}
	key := entity.Key(k)
	val := entity.Value{Val: v, Exp: exp}

	s.Lock()
	defer s.Unlock()
	s.m[key] = val
	return nil
}

func (s *Store) Delete(_ context.Context, k string) error {
	s.safeRemove(entity.Key(k))
	return nil
}

func (s *Store) Clean(_ context.Context) error {
	// Cleaning up process
	for idx, key := range s.m {
		if !key.IsValid() {
			s.safeRemove(idx)
		}
	}
	return nil
}

func (s *Store) Shot(_ context.Context, path string) error {
	fmode := 0644
	s.Lock()
	defer s.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	marsh, err := json.Marshal(s.m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, marsh, os.FileMode(fmode))
}

func NewStore() Store {
	m := make(map[entity.Key]entity.Value)
	return Store{
		m:       m,
		RWMutex: sync.RWMutex{},
	}
}

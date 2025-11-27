package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ICache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
	DeleteAll() error
}

type LocalCache struct {
	items    map[string]any
	ttl      time.Time // for evection
	capacity int

	sync.Mutex
	// TODO: Eviction
}

func NewLocalCache(cap int) (*LocalCache, error) {
	if cap <= 0 {
		return nil, errors.New("capacity must be greater than zero")
	}

	return &LocalCache{
		items:    make(map[string]any, cap),
		ttl:      time.Now(),
		capacity: cap,
	}, nil
}

func (c *LocalCache) Set(key string, value any) error {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value

	return nil
}

func (c *LocalCache) Get(key string) (any, error) {
	c.Lock()
	defer c.Unlock()

	value, ok := c.items[key]
	if !ok {
		return nil, errors.New("key does not exists")
	}
	return value, nil
}

func (c *LocalCache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	_, ok := c.items[key]
	if !ok {
		return errors.New("key does not exists")
	}

	delete(c.items, key)

	return nil
}

func (c *LocalCache) DeleteAll() error {
	c.Lock()
	defer c.Unlock()

	for key := range c.items {
		delete(c.items, key)
	}

	return nil
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	lc, err := NewLocalCache(5)
	if err != nil {
		panic(err)
	}

	lc.Set("firstName", "fuzail ahmed")
	lc.Set("lastName", "Ansari")
	lc.Set("age", 32)

	name, _ := lc.Get("firstName")
	fmt.Printf("%v\n", name)

	lname, _ := lc.Get("lastName")
	fmt.Printf("%v\n", lname)

	lc.Delete("firstName")

	name2, err := lc.Get("firstName")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", name2)
}

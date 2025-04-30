package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	defaultDuration     = 1 * time.Minute
	expiredLoopDuration = 1 * time.Second
)

type Cache struct {
	duration time.Duration
	items    map[string]cacheItem
	mutex    sync.RWMutex
	done     chan struct{}
	once     sync.Once
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCache(duration *time.Duration) *Cache {
	c := &Cache{
		items: make(map[string]cacheItem),
		done:  make(chan struct{}),
	}

	// use default expiration time if none is passed when creating the cache
	if duration != nil {
		c.duration = *duration
	} else {
		c.duration = defaultDuration
	}

	// keep the clean up going
	go c.expiredItemsLoop()

	return c
}

func (c *Cache) Add(key string, value interface{}, duration *time.Duration) {
	// use cache duration if one was not passed to the item
	if duration == nil || (duration != nil && *duration <= 0) {
		duration = &c.duration
	}

	expiration := time.Now().Add(*duration)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// add to cache
	c.items[key] = cacheItem{
		value:      value,
		expiration: expiration,
	}

	fmt.Println("Item added to cache", key)
}

func (c *Cache) Get(key string) (interface{}, error) {
	// for thread safety
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, errors.New("key not found")
	}

	if time.Now().After(item.expiration) {
		return nil, errors.New("key expired")
	}

	return item.value, nil
}

func (c *Cache) Delete(key string) {
	// for thread safety
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)

	fmt.Println("Item deleted from cache", key)
}

func (c *Cache) Close() {
	c.once.Do(func() {
		close(c.done)
	})
}

func (c *Cache) expiredItemsLoop() {
	// run periodically based on the default expiration duration
	ticker := time.NewTicker(expiredLoopDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.removeExpiredItems()
		case <-c.done:
			return
		}
	}
}

func (c *Cache) removeExpiredItems() {
	// for thread safety
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// traverse the items in the cache and delete all the ones which expiration
	// are past the current time
	for k, v := range c.items {
		if v.expiration.After(time.Now()) {
			delete(c.items, k)

			fmt.Println("Item expired", k)
		}
	}
}

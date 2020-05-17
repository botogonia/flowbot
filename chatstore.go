package flowbot

import (
	"sync"
)

type ChatStore struct {
	mx sync.RWMutex
	m  map[int64]*Chat
}

func (c *ChatStore) Get(key int64) (*Chat, bool) {
	c.mx.RLock()
	val, ok := c.m[key]
	c.mx.RUnlock()
	return val, ok
}

func (c *ChatStore) Save(key int64, value *Chat) {
	c.mx.Lock()
	c.m[key] = value
	c.mx.Unlock()
}

func (c *ChatStore) Del(key int64) {
	c.mx.Lock()
	delete(c.m, key)
	c.mx.Unlock()
}

func NewChatStore() *ChatStore {
	return &ChatStore{
		m: make(map[int64]*Chat),
	}
}

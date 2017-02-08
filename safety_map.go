package mdzz

import (
	"sync"
)

type Safetymap struct {
	sync.RWMutex
	items map[string]interface{}
}

func NewSafetyMap() *Safetymap {
	return &Safetymap{
		items: make(map[string]interface{}),
	}
}

func (m *Safetymap) Get(key string) (value interface{}, ok bool) {
	m.RLock()
	value, ok = m.items[key]
	m.RUnlock()
	return
}

func (m *Safetymap) Set(key string, value interface{}) {
	m.Lock()
	m.items[key] = value
	m.Unlock()
}

func (m *Safetymap) Delete(key string) {
	m.Lock()
	delete(m.items, key)
	m.Unlock()
}

func (m *Safetymap) Has(key string) (ok bool) {
	m.RLock()
	_, ok = m.items[key]
	m.RUnlock()
	return
}

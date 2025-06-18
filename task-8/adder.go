package task8

import (
	"fmt"
	"sync"
)

type MyMap struct {
	list map[int]int
	mx sync.RWMutex
}

func New() *MyMap {
	return &MyMap{
		list: make(map[int]int),
	}
}

func (m *MyMap) InsertValue(x int, y int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.list[x] = y
}

func (m *MyMap) ReadValue(x int) (int, error) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	res, ok := m.list[x]
	if ok {
		return res, nil
	}

	return 0, fmt.Errorf("not found") 
}
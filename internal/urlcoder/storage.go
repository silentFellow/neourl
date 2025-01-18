package urlcoder

import "sync"

type Storage struct {
	lookup        map[string]string
	reverseLookup map[string]string

	charLength   int
	mapLength    int
	maxMapLength int

	lock sync.RWMutex
}

func NewURLStorage() *Storage {
	return &Storage{
		lookup:        make(map[string]string),
		reverseLookup: make(map[string]string),

		charLength:   1,
		mapLength:    0,
		maxMapLength: 31,

		lock: sync.RWMutex{},
	}
}

func (s *Storage) findNextMaxCharLength() int {
	newMaxMapLength := 1
	for range s.charLength + 1 {
		newMaxMapLength *= 62
	}
	return newMaxMapLength / 2
}

package urlcoder

import "sync"

type Storage struct {
	lookup        map[string]string
	reverseLookup map[string]string

	mapLength     int
	charLength    int
	maxCharLength int

	lock sync.RWMutex
}

func NewURLStorage() *Storage {
	return &Storage{
		lookup:        make(map[string]string),
		reverseLookup: make(map[string]string),

		charLength:    1,
		maxCharLength: 31,

		lock: sync.RWMutex{},
	}
}

func (s *Storage) findNextMaxCharLength() int {
	return (s.maxCharLength * 62) / 2
}

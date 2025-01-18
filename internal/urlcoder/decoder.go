package urlcoder

import "errors"

func (s *Storage) DecodeURL(encoded string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	val, ok := s.lookup[encoded]
	if !ok {
		return "", errors.New("Url not found")
	}

	return val, nil
}

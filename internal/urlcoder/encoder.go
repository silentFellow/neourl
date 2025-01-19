package urlcoder

import (
	"math/rand/v2"
	"strings"
)

func (s *Storage) EncodeURL(url string) string {
	url = formatURL(url)

	s.lock.RLock()
	val, ok := s.reverseLookup[url]
	s.lock.RUnlock()

	if ok {
		return val
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.mapLength+1 > s.maxMapLength {
		s.charLength += 1
		s.maxMapLength = s.findNextMaxCharLength()
	}

	encoded := generateRandom(s.charLength)
	for {
		if _, ok := s.reverseLookup[encoded]; !ok {
			break
		}
		encoded = generateRandom(s.charLength)
	}
	s.lookup[encoded] = url
	s.reverseLookup[url] = encoded
	s.mapLength += 1

	return encoded
}

func generateRandom(length int) string {
	var encoded strings.Builder

	for range length {
		valType := rand.IntN(3)

		switch valType {
		case 0:
			encoded.WriteByte(byte('a' + rand.IntN(26)))
		case 1:
			encoded.WriteByte(byte('0' + rand.IntN(10)))
		case 2:
			encoded.WriteByte(byte('A' + rand.IntN(26)))
		}
	}

	return encoded.String()
}

func formatURL(url string) string {
	neoURL := strings.ReplaceAll(url, " ", "")
	neoURL = strings.TrimSuffix(neoURL, "/")

	if !strings.HasPrefix(neoURL, "http") && !strings.HasPrefix(neoURL, "https") {
		neoURL = "http://" + neoURL
	}

	return neoURL
}

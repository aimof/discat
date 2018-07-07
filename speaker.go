package catheta

import (
	"math/rand"
	"strings"
)

type speaker struct {
	table map[string][]string
}

func newSpeaker(input [][]string) speaker {
	s := speaker{
		table: make(map[string][]string, len(input)),
	}
	for _, words := range input {
		if len(words) < 2 {
			continue
		}
		s.table[words[0]] = words[1:]
	}
	return s
}

func (s speaker) speak(sentenceIn string) string {
	for key, value := range s.table {
		if strings.Contains(sentenceIn, key) {
			return value[rand.Intn(len(value))]
		}
	}
	return ""
}

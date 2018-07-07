package discat

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
		empty := len(words)
		for i := 1; i < len(words); i++ {
			if words[i] == "" {
				empty = i
				break
			}
		}
		s.table[words[0]] = words[1:empty]
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

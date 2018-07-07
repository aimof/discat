package discat

import (
	"reflect"
	"testing"
)

func TestNewSpeaker(t *testing.T) {
	testCases := []struct {
		input    [][]string
		expected speaker
	}{
		{
			[][]string{
				{"Hello", "Greeting", "Hello, world!"},
				{"GoodBye", "Greeting"},
			},
			speaker{
				table: map[string][]string{
					"Hello":   {"Greeting", "Hello, world!"},
					"GoodBye": {"Greeting"},
				},
			},
		},
	}
	for i, tt := range testCases {
		s := newSpeaker(tt.input)
		if !reflect.DeepEqual(s, tt.expected) {
			t.Errorf("case: %d\nexpected: %v\nactual: %v", i, tt.expected, s)
		}
	}
}

func TestSpeak(t *testing.T) {
	s := newSpeaker([][]string{
		{"Hello", "Greeting!"},
		{"こんにちは", "Greeting!", "よう！"},
	})
	testCases := map[string][]string{
		"Hello, world!": {"Greeting!"},
		"こんにちは、お元気ですか？": {"Greeting!", "よう！"},
		"Go my way.":    {""},
	}
test1case:
	for key, value := range testCases {
		actual := s.speak(key)
		for _, expected := range value {
			if actual == expected {
				break test1case
			}
		}
		t.Errorf("case: %s, expected: %s, actual: %s", key, value, actual)
	}
}

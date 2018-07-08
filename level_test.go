package discat

import "testing"

func TestExpUp(t *testing.T) {
	l := level{exp: make(map[string]int)}
	l.gainExp("name", 1)
	if l.exp["name"] != 1 {
		t.Error()
	}
	l.gainExp("name", 2)
	if l.exp["name"] != 3 {
		t.Error()
	}
}

func TestshowLevel(t *testing.T) {
	testCases := []struct {
		name          string
		currentExp    int
		expectedLevel int
		expectedRest  int
	}{
		{"case0", 99, 12, 10},
		{"case1", 0, 0, 3},
		{"case2", 5, 1, 4},
	}
	for _, tt := range testCases {
		l := level{exp: make(map[string]int)}
		l.exp[tt.name] = tt.currentExp
		exp, level, rest := l.showLevel(tt.name)
		if exp != tt.currentExp || level != tt.expectedLevel || rest != tt.expectedRest {
			t.Errorf("case: %s\ncurrent exp: %d\nexpected: %d, %d\nactual: %d,%d", tt.name, tt.currentExp, tt.expectedLevel, tt.expectedRest, level, rest)
		}
	}
}

func TestExpToLevel(t *testing.T) {
	testCases := []struct {
		exp      int
		expected int
	}{
		{99, 12},
		{0, 0},
		{5, 1},
	}
	for _, tt := range testCases {
		l := expToLevel(tt.exp)
		if l != tt.expected {
			t.Errorf("Input: %d, expected:%d, actual: %d", tt.exp, tt.expected, l)
		}
	}
}

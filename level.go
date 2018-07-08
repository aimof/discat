package discat

import "math"

type level struct {
	exp map[string]int
}

func (l level) gainExp(name string, n int) {
	if _, ok := l.exp[name]; ok {
		l.exp[name] += n
	} else {
		l.exp[name] = n
	}
}

func (l level) showLevel(name string) (exp, level, rest int) {
	exp = l.exp[name]
	level = expToLevel(l.exp[name])
	// safety: 100000
	for i := l.exp[name] + 1; i < l.exp[name]+100000; i++ {
		if j := expToLevel(i); j > level {
			return exp, level, i - l.exp[name]
		}

	}
	return 0, 0, 0
}

func expToLevel(exp int) int {
	return int(math.Log10(float64(exp)+1) + (float64(exp)+1)/10)
}

func (l level) output() map[string]int { return l.exp }

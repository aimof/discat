package discat

type Discat struct {
	speaker speaker
	level   level
}

func Init(dict [][]string, expMap map[string]int) Discat {
	return Discat{
		speaker: newSpeaker(dict),
		level:   level{exp: expMap},
	}
}

func (d Discat) Speak(input string) string {
	return d.speaker.speak(input)
}

func (d Discat) GainExp(name string, exp int) {
	d.level.gainExp(name, exp)
}

func (d Discat) ShowLevel(name string) (exp, level, rest int) {
	return d.level.showLevel(name)
}

func (d Discat) Output() map[string]int { return d.level.output() }

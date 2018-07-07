package discat

import "github.com/bwmarrin/discordgo"

type Discat struct {
	speaker speaker
}

func Init(path string) (Discat, error) {
	s := newStarter(path)
	dict, err := s.readDictionary()
	if err != nil {
		return Discat{}, err
	}
	return Discat{
		speaker: newSpeaker(dict),
	}, err
}

func (d Discat) Speak(m *discordgo.MessageCreate) string {
	return d.speaker.speak(m.Content)
}

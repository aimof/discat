package discat

import "github.com/bwmarrin/discordgo"

type Discat struct {
	speaker speaker
}

func Init(dict [][]string) Discat {
	return Discat{
		speaker: newSpeaker(dict),
	}
}

func (d Discat) Speak(m *discordgo.MessageCreate) string {
	return d.speaker.speak(m.Content)
}

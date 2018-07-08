package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aimof/discat"
	"github.com/bwmarrin/discordgo"
)

var (
	token    string
	dictPath string
	name     string
	nickname string
	dcat     discat.Discat
)

func main() {
	err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	s := newStarter(dictPath)
	dict, err := s.readDictionary()
	if err != nil {
		log.Fatalln(err)
	}

	dcat = discat.Init(dict)
	if err != nil {
		log.Fatalln(err)
	}

	session, err := discordgo.New()
	if err != nil {
		log.Fatalln(err)
	}
	session.Token = token

	session.AddHandler(handler)

	err = session.Open()
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	_ = <-ch
	err = session.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Printf("\t%v\n", err)
	}

	var msg string
	if m.Author.Username != name {
		log.Printf("\tMessage from %s: %s\n", m.Author.Username, m.Content)
		msg = dcat.Speak(m)
	}

	if msg != "" {
		_, err = s.ChannelMessageSend(c.ID, msg)
		if err != nil {
			log.Printf("\t%v\n", err)
		}
		log.Printf("\tMessage sent: %s\n", msg)
	}
}

func readConfig() error {
	file, err := os.Open("config.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		return err
	}

	var config []string
	for scanner.Scan() {
		config = append(config, scanner.Text())
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	token = config[0]
	dictPath = config[1]
	name = config[2]
	nickname = config[3]
	return nil
}

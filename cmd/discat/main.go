package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/aimof/discat"
	"github.com/bwmarrin/discordgo"
)

var (
	name     string
	nickname string
	dcat     discat.Discat
)

func main() {
	token, dictPath, err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	s := newStarter(dictPath, "data/level.csv")
	dict, err := s.readDictionary()
	if err != nil {
		log.Fatalln(err)
	}

	levelMap, err := s.readLevel()
	if err != nil {
		log.Fatalln(err)
	}

	dcat = discat.Init(dict, levelMap)
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

	levelMap = dcat.Output()
	s.writeLevel(levelMap)
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Printf("\t%v\n", err)
	}

	dcat.GainExp(m.Author.Username, 1)

	var msg string
	if m.Author.Username != name {
		log.Printf("\tMessage from %s: %s\n", m.Author.Username, m.Content)
		if strings.Contains(m.Content, name) || strings.Contains(m.Content, nickname) {
			if strings.Contains(m.Content, "!level") {
				e, l, r := dcat.ShowLevel(m.Author.Username)
				msg = "Exp: " + strconv.Itoa(e) + ",Level: " + strconv.Itoa(l) + ",Next to: " + strconv.Itoa(r)
			}
		}
		if msg == "" {
			msg = dcat.Speak(m.Content)
		}
	}

	if msg != "" {
		_, err = s.ChannelMessageSend(c.ID, msg)
		if err != nil {
			log.Printf("\t%v\n", err)
		}
		log.Printf("\tMessage sent: %s\n", msg)
	}
}

func readConfig() (token, dictPath string, err error) {
	file, err := os.Open("config.txt")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		return "", "", err
	}

	var config []string
	for scanner.Scan() {
		config = append(config, scanner.Text())
	}

	if scanner.Err() != nil {
		return "", "", scanner.Err()
	}

	name = config[1]
	nickname = config[2]
	return config[0], config[3], nil
}

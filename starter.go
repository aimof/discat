package discat

import (
	"encoding/csv"
	"os"
)

type starter struct {
	configPath string
}

func newStarter(configPath string) starter {
	return starter{configPath: configPath}
}

func (s starter) readDictionary() ([][]string, error) {
	f, err := os.Open(s.configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	return reader.ReadAll()
}

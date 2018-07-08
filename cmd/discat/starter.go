package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
)

type starter struct {
	configPath string
	levelPath  string
}

func newStarter(configPath, levelPath string) starter {
	return starter{configPath: configPath, levelPath: levelPath}
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

func (s starter) readLevel() (levelMap map[string]int, err error) {
	_, err = os.Stat(s.levelPath)
	if err != nil {
		return make(map[string]int, 60), nil
	}

	f, err := os.Open(s.levelPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	levelTable, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return tableToMap(levelTable)
}

func tableToMap(levelTable [][]string) (levelMap map[string]int, err error) {
	levelMap = make(map[string]int, len(levelTable))
	for _, line := range levelTable {
		if len(line) != 2 {
			return nil, errors.New("level.csv's rows are not 2")
		}

		n, err := strconv.Atoi(line[1])
		if err != nil {
			log.Printf("Error: level.csv contains non-int valule\n")
			n = 0
		}
		levelMap[line[0]] = n
	}
	return levelMap, nil

}

func (s starter) writeLevel(levelMap map[string]int) error {
	f, err := os.OpenFile(s.levelPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	levelTable := mapToTable(levelMap)

	writer := csv.NewWriter(f)
	return writer.WriteAll(levelTable)
}

func mapToTable(levelMap map[string]int) (levelTable [][]string) {
	levelTable = make([][]string, 0, len(levelMap))
	for key, value := range levelMap {
		levelTable = append(levelTable, []string{key, strconv.Itoa(value)})
	}
	return sortTable(levelTable)
}

type table [][]string

func (t table) Len() int           { return len(t) }
func (t table) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t table) Less(i, j int) bool { return t[i][0] < t[j][0] }

func sortTable(t [][]string) [][]string {
	tbl := table(t)
	sort.Sort(tbl)
	return [][]string(tbl)
}

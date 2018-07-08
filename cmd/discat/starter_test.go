package main

import (
	"reflect"
	"testing"
)

func TestReadDictionary(t *testing.T) {
	testCases := []struct {
		configPath  string
		expected    [][]string
		expectedErr error
	}{
		{
			configPath: "test/TestCase0.csv",
			expected: [][]string{
				[]string{"0", "1"},
				[]string{"1", "2"},
			},
			expectedErr: nil,
		},
		{
			configPath: "test/TestCase1.csv",
			expected: [][]string{
				[]string{"Hello", "Greeting"},
				[]string{"こんにちは", "Greeting"},
			},
			expectedErr: nil,
		},
	}
	for i, tt := range testCases {
		s := newStarter(tt.configPath)
		actual, actualErr := s.readDictionary()
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("case: %d\nexpected: %v\nactual%v", i, tt.expected, actual)
		}
		if !reflect.DeepEqual(actualErr, tt.expectedErr) {
			t.Errorf("case: %d\nexpected: %v\nactual%v", i, tt.expectedErr, actualErr)
		}
	}
}

var mapAndTableCases = []struct {
	levelMap   map[string]int
	levelTable [][]string
	err        error
}{
	{
		levelMap: map[string]int{
			"case0": 0,
			"case1": 1,
			"case2": 16383,
			"case3": -1,
		},
		levelTable: [][]string{
			{"case0", "0"},
			{"case1", "1"},
			{"case2", "16383"},
			{"case3", "-1"},
		},
	},
}

func TestTableToMap(t *testing.T) {
	for _, tt := range mapAndTableCases {
		m, err := tableToMap(tt.levelTable)
		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("expected: %v\n, actual%v", tt.err, err)
		}
		if !reflect.DeepEqual(m, tt.levelMap) {
			t.Errorf("expected: %v\nactual: %v", tt.levelMap, m)
		}
	}
}

func TestMapToTable(t *testing.T) {
	for _, tt := range mapAndTableCases {
		if tt.err != nil {
			continue
		}
		tbl := sortTable(mapToTable(tt.levelMap))
		if !reflect.DeepEqual(tbl, tt.levelTable) {
			t.Errorf("expected: %v\nactual: %v", tt.levelTable, tbl)
		}
	}
}

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

package command

import (
	"reflect"
	"testing"
)

func TestParseCommandArgs(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected rawCommand
	}{
		{"WithoutArg", "/teste", rawCommand{"teste", []string{}}},
		{"WithOneArg", "/teste 1", rawCommand{"teste", []string{"1"}}},
		{"WithManyArgs", "/teste 1 2 3", rawCommand{"teste", []string{"1", "2", "3"}}},
		{"WithTextArgs", "/teste one two three", rawCommand{"teste", []string{"one", "two", "three"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseCommandArgs(tc.input)

			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, got)
			}
		})
	}
}

func TestParseArgHandler(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected ArgHandler
	}{
		{"WithoutArg", []string{}, ArgHandler{}},
		{"WithOneArg", []string{"1"}, ArgHandler{"1", []string{"1"}, 0}},
		{"WithManyArgs", []string{"1", "2", "3"}, ArgHandler{"1 2 3", []string{"1", "2", "3"}, 0}},
		{"WithTextArgs", []string{"one", "two", "three"}, ArgHandler{"one two three", []string{"one", "two", "three"}, 0}},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseArgHandler(tc.input)

			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, got)
			}
		})
	}
}

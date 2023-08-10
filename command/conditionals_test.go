package command

import "testing"
import "reflect"

func TestCreateConditional(t *testing.T) {
	t.Skip()
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

func TestBuilderConditional(t *testing.T) {
	type result struct {
		idx     int
		typeIdx int
	}

	testCases := []struct {
		name     string
		idx      int
		typeIdx  int
		expected result
	}{
		{name: "test index", idx: 0, expected: result{idx: 0}},
		{name: "test index", idx: 0, expected: result{idx: 0}},
		{name: "test type index", typeIdx: 0, expected: result{typeIdx: 0}},
		{name: "test type index", typeIdx: 0, expected: result{typeIdx: 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := Builder()
			cmd.Conditionals().
				Index(tc.idx)

			switch tc.typeIdx {
			case typePlayer:
				cmd.Conditionals().TypePlayer()
			case typeNumber:
				// comentado: ainda não está pronto
				// cmd.Conditionals().TypeNumber()
			case typeText:
				// comentado: ainda não está pronto
				// cmd.Conditionals().TypeText
			}

			got := cmd.Conditionals().index
			if !reflect.DeepEqual(got, tc.expected.idx) {
				t.Errorf("expeted: %d; got: %d", tc.expected.idx, got)
			}

			got = cmd.Conditionals().typeIdx
			if !reflect.DeepEqual(got, tc.expected.typeIdx) {
			}
		})
	}
}

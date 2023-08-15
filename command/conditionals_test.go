package command

import (
	"reflect"
	"testing"
)

func TestCreateConditional(t *testing.T) {
	testCases := []struct {
		name          string
		inputTypeCond int
		inputTypeIdx  int
		inputValue    interface{}
		expected      []condition
	}{
		{
			"PlayerCondition",
			0,
			typePlayer,
			nil,
			[]condition{{typePlayer, 0, nil}},
		},
		{
			"NumberCondition",
			0,
			typeNumber,
			nil,
			[]condition{{typeNumber, 0, nil}},
		},
		{
			"TextCondition",
			0,
			typeText,
			nil,
			[]condition{{typeText, 0, nil}},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cb := Builder().
				Conditionals()

			cb.createConditional(tc.inputTypeCond, tc.inputTypeIdx, tc.inputValue)

			got := cb.conditions

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
		{name: "type undefined index", idx: 0, typeIdx: -1, expected: result{0, 0}},
		{name: "type undefined index", idx: 0, typeIdx: 0, expected: result{0, 0}},
		{name: "type player index", idx: 0, typeIdx: typePlayer, expected: result{0, typePlayer}},
		{name: "type text index", idx: 0, typeIdx: typeText, expected: result{0, typeText}},
		{name: "type number index", idx: 0, typeIdx: typeNumber, expected: result{0, typeNumber}},

		{name: "index", idx: -1, typeIdx: 0, expected: result{0, 0}},
		{name: "index", idx: 0, typeIdx: 0, expected: result{0, 0}},
		{name: "index", idx: 10, typeIdx: 0, expected: result{10, 0}},
		{name: "index", idx: 30, typeIdx: 0, expected: result{30, 0}},
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
				cmd.Conditionals().TypeNumber()
			case typeText:
				// comentado: ainda não está pronto
				cmd.Conditionals().TypeText()
			}

			got := cmd.Conditionals().index
			if !reflect.DeepEqual(got, tc.expected.idx) {
				t.Errorf("expeted idx: %d; got: %d", tc.expected.idx, got)
			}

			got = cmd.Conditionals().typeIdx
			if !reflect.DeepEqual(got, tc.expected.typeIdx) {
				t.Errorf("expeted type: %d; got: %d", tc.expected.typeIdx, got)
			}
		})
	}
}

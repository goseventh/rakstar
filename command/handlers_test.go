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

func TestStrFuncsTypeNumber(t *testing.T) {
	type result struct {
		value bool
	}
	testCases := []struct {
		name      string
		typeTest  uint8
		arg       string
		condValue interface{}
		expected  result
	}{
		{"valueStrEqual", MustEqual, "30", 30, result{value: true}},
		{"valueStrEqual", MustEqual, "30", 10, result{value: false}},
		{"valueStrBeBetween", MustBeBetween, "7", []int{10, 5}, result{value: true}},
		{"valueStrBeBetween", MustBeBetween, "5", []int{3, 10}, result{value: false}},
		{"valueStrBeLess", MustBeLessThan, "3", 7, result{value: true}},
		{"valueStrBeLess", MustBeLessThan, "7", 7, result{value: false}},
		{"valueStrBeGreeter", MustBeGreaterThan, "120", 30, result{value: true}},
		{"valueStrBeGreeter", MustBeGreaterThan, "30", 50, result{value: false}},
		{"valueStrSquareRootOf", MustBeSquareRootOf, "2", 4, result{value: true}},
		{"valueStrSquareRootOf", MustBeSquareRootOf, "2", 16, result{value: false}},
		{"valueStrDivisibleBy", MustBeDivisibleBy, "10", 2, result{value: true}},
		{"valueStrDivisibleBy", MustBeDivisibleBy, "50", 3, result{value: false}},
    {"valueStrMultipleOfBy", MustBeMultipleOf, "10", 2, result{value: true}},
		{"valueStrMultipleOfBy", MustBeMultipleOf, "2", 10, result{value: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.typeTest {
			case MustEqual:
				got := valueStrEqual(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeBetween:
				got := valueStrBeBetween(tc.arg, tc.condValue.([]int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeGreaterThan:
				got := valueStrBeGreeter(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeLessThan:
				got := valueStrBeLess(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeMultipleOf:
				got := valueStrDivisibleBy(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeSquareRootOf:
				got := valueStrSquareRootOf(tc.arg, tc.condValue.(int))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			}
		})
	}
}

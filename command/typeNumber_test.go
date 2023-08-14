package command

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goseventh/rakstar/internal/natives"
)

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
		{"valueStrEqualFail", MustEqual, "30", 10, result{value: false}},
		{"valueStrBeBetween", MustBeBetween, "7", []int{10, 5}, result{value: true}},
		{"valueStrBeBetweenFail", MustBeBetween, "5", []int{3, 10}, result{value: false}},
		{"valueStrBeLess", MustBeLessThan, "3", 7, result{value: true}},
		{"valueStrBeLessFail", MustBeLessThan, "7", 7, result{value: false}},
		{"valueStrBeGreeter", MustBeGreaterThan, "120", 30, result{value: true}},
		{"valueStrBeGreeterFail", MustBeGreaterThan, "30", 50, result{value: false}},
		{"valueStrSquareRootOf", MustBeSquareRootOf, "2", 4, result{value: true}},
		{"valueStrSquareRootOfFail", MustBeSquareRootOf, "2", 16, result{value: false}},
		{"valueStrDivisibleBy", MustBeDivisibleBy, "10", 2, result{value: true}},
		{"valueStrDivisibleByFail", MustBeDivisibleBy, "50", 3, result{value: false}},
		{"valueStrMultipleOfBy", MustBeMultipleOf, "10", 2, result{value: true}},
		{"valueStrMultipleOfByFail", MustBeMultipleOf, "2", 10, result{value: false}},
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

func TestCondicionalsNumbers(t *testing.T) {
	commands = make(map[string]*Command)
	type result struct {
		value bool
	}
	testCases := []struct {
		name      string
		typeTest  uint8
		index     int
		args      string
		condValue interface{}
		expected  result
	}{
		{
			name:      "cmdMustEqual",
			typeTest:  MustEqual,
			index:     0,
			args:      "30",
			condValue: 30,
			expected:  result{value: true}},
		{
			name:      "cmdMustBeBetween",
			typeTest:  MustBeBetween,
			index:     0,
			args:      "7",
			condValue: []int{10, 5},
			expected:  result{value: true}},
		{
			name:      "cmdMustBeLess",
			typeTest:  MustBeLessThan,
			index:     0,
			args:      "3",
			condValue: 7,
			expected:  result{value: true}},
		{
			name:      "cmdMustBeGreeter",
			typeTest:  MustBeGreaterThan,
			index:     0,
			args:      "120",
			condValue: 30,
			expected:  result{value: true}},
		{
			name:      "cmdMustBeSquareRootOf",
			typeTest:  MustBeSquareRootOf,
			index:     0,
			args:      "2",
			condValue: 4,
			expected:  result{value: true}},
		{
			name:      "cmdMustBeDivisibleBy",
			typeTest:  MustBeDivisibleBy,
			index:     0,
			args:      "10",
			condValue: 2,
			expected:  result{value: true}},
		{
			name:      "cmdMustBeMultipleOfBy",
			typeTest:  MustBeMultipleOf,
			index:     0,
			args:      "10",
			condValue: 2,
			expected:  result{value: true}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.typeTest {
			case MustEqual:
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustEqual(tc.condValue.(int)).
					EndConditionals().
					Create()
				t.Logf("COMAND:%v", fmt.Sprintf("/%v %v", tc.name, tc.args))
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeBetween:
				params := tc.condValue.([]int)
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustBeBetween(params[0], params[1]).
					EndConditionals().
					Create()
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeGreaterThan:
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustBeGreaterThan(tc.condValue.(int)).
					EndConditionals().
					Create()
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeLessThan:
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustBeLessThan(tc.condValue.(int)).
					EndConditionals().
					Create()
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeMultipleOf:
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustBeMultipleOf(tc.condValue.(int)).
					EndConditionals().
					Create()
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeSquareRootOf:
				Builder().
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeNumber().
					MustBeSquareRootOf(tc.condValue.(int)).
					EndConditionals().
					Create()
				got := processCommand(
					natives.Player{}, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			}
		})
	}
}

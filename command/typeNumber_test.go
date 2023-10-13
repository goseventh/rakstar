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
		{"valueStrEqualFail", MustEqual, "130", 30, result{value: false}},
<<<<<<< HEAD
		{"valueStrBeBetween", MustBeBetween, "7", []int{10, 5}, result{value: true}},
		{"valueStrBeBetweenFail", MustBeBetween, "5", []int{3, 10}, result{value: false}},
=======
		{"valueStrBeBetween", MustBeBetween, "7", []int{5, 10}, result{value: true}},
		{"valueStrBeBetweenFail", MustBeBetween, "5", []int{10, 3}, result{value: false}},
>>>>>>> dev
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
<<<<<<< HEAD
				got := valueStrBeBetween(tc.arg, tc.condValue.([]int))
=======
				got := valueStrBeBetween(tc.arg, tc.condValue.([]int)[0], tc.condValue.([]int)[1])
>>>>>>> dev
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
		executeCommand bool
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
			expected:  result{executeCommand: true}},
		{
			name:      "cmdMustEqualFail",
			typeTest:  MustEqual,
			index:     0,
			args:      "5",
			condValue: 2,
			expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeBetween",
			typeTest:  MustBeBetween,
			index:     7,
			args:      "7",
			condValue: []int{10, 5},
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeBetweenFail",
		// 	typeTest:  MustBeBetween,
		// 	index:     5,
		// 	args:      "15",
		// 	condValue: []int{10, 5},
		// 	expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeLess",
			typeTest:  MustBeLessThan,
			index:     3,
			args:      "3",
			condValue: 7,
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeLessFail",
		// 	typeTest:  MustBeLessThan,
		// 	index:     3,
		// 	args:      "3",
		// 	condValue: 0,
		// 	expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeGreeter",
			typeTest:  MustBeGreaterThan,
			index:     2,
			args:      "120",
			condValue: 30,
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeGreeterFail",
		// 	typeTest:  MustBeGreaterThan,
		// 	index:     -1,
		// 	args:      "40",
		// 	condValue: 90,
		// 	expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeSquareRootOf",
			typeTest:  MustBeSquareRootOf,
			index:     -1,
			args:      "2",
			condValue: 4,
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeSquareRootOfFail",
		// 	typeTest:  MustBeSquareRootOf,
		// 	index:     3,
		// 	args:      "6",
		// 	condValue: 4,
		// 	expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeDivisibleBy",
			typeTest:  MustBeDivisibleBy,
			index:     1,
			args:      "10",
			condValue: 2,
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeDivisibleByFail",
		// 	typeTest:  MustBeDivisibleBy,
		// 	index:     0,
		// 	args:      "100",
		// 	condValue: 200,
		// 	expected:  result{executeCommand: false}},
		{
			name:      "cmdMustBeMultipleOfBy",
			typeTest:  MustBeMultipleOf,
			index:     8,
			args:      "10",
			condValue: 2,
			expected:  result{executeCommand: true}},
		// {
		// 	name:      "cmdMustBeMultipleOfByFail",
		// 	typeTest:  MustBeMultipleOf,
		// 	index:     0,
		// 	args:      "34",
		// 	condValue: 25,
		// 	expected:  result{executeCommand: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := Builder()
			cmd.Conditionals().Index(tc.index)
			switch tc.typeTest {
			case MustEqual:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustEqual(tc.condValue.(int)).
					EndConditionals().
					Create()
				t.Logf("COMAND:%v", fmt.Sprintf("/%v %v", tc.name, tc.args))
				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			case MustBeBetween:
				params := tc.condValue.([]int)
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustBeBetween(params[0], params[1]).
					EndConditionals().
					Create()
				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			case MustBeGreaterThan:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustBeGreaterThan(tc.condValue.(int)).
					EndConditionals().
					Create()
				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			case MustBeLessThan:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustBeLessThan(tc.condValue.(int)).
					EndConditionals().
					Create()

				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			case MustBeMultipleOf:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustBeMultipleOf(tc.condValue.(int)).
					EndConditionals().
					Create()
				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			case MustBeSquareRootOf:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					// Index(tc.index).
					TypeNumber().
					MustBeSquareRootOf(tc.condValue.(int)).
					EndConditionals().
					Create()
				player := new(natives.Player)
				got := processCommand(
					*player, fmt.Sprintf("/%v %v", tc.name, tc.args),
				)
				if !reflect.DeepEqual(got, tc.expected.executeCommand) {
					t.Fatalf("expected: %v; got: %v", tc.expected.executeCommand, got)
				}
			}
		})
	}
}

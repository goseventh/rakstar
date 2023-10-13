package command

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goseventh/rakstar/internal/natives"
)

func TestTypeText(t *testing.T) {
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
		{"textIsUpper", MustBeUppercase, "TEST", nil, result{value: true}},
		{"textIsUpperFail", MustBeUppercase, "test", nil, result{value: false}},
		{"textIsUpperFailCapital", MustBeUppercase, "Test", nil, result{value: false}},

		{"textIsLower", MustBeLowercase, "test", nil, result{value: true}},
		{"textIsLowerFail", MustBeLowercase, "TEST", nil, result{value: false}},
		{"textIsLowerFailCapital", MustBeLowercase, "Test", nil, result{value: false}},

		{"textIsPrefix", MustHavePrefix, "testrakstar", "test", result{value: true}},
		{"textIsPrefixFail", MustHavePrefix, "tastrakstar", "test", result{value: false}},
		{"textIsPrefixFailUpper", MustHavePrefix, "TESTRAKSTAR", "test", result{value: false}},

		{"textIsSuffix", MustHaveSufix, "rakstartest", "test", result{value: true}},
		{"textIsSuffixFail", MustHaveSufix, "rakstartast", "test", result{value: false}},
		{"textIsSuffixFailUpper", MustHaveSufix, "RAKSTARTEST", "test", result{value: false}},

		{"textIsRegMatch", MustCompileRegex, "support@rakstar.com", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, result{value: true}},
		{"textIsRegMatchFail", MustCompileRegex, "support@rakstar", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, result{value: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.typeTest {
			case MustBeUppercase:
				got := textIsUpper(tc.arg)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustBeLowercase:
				got := textIsLower(tc.arg)
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustHavePrefix:
				got := textIsPrefix(tc.arg, tc.condValue.(string))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustHaveSufix:
				got := textIsSuffix(tc.arg, tc.condValue.(string))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			case MustCompileRegex:
				got := textIsRegMatch(tc.arg, tc.condValue.(string))
				if !reflect.DeepEqual(got, tc.expected.value) {
					t.Fatalf("expected: %v; got: %v", tc.expected.value, got)
				}
			}
		})
	}
}

func TestCondicionalsTexts(t *testing.T) {
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
		{"cmdMustBeUppercase", MustBeUppercase, 0, "TEST", nil, result{executeCommand: true}},
		{"cmdMustBeUppercaseFail", MustBeUppercase, 0, "test", nil, result{executeCommand: false}},
		{"cmdMustBeUppercaseFailCapital", MustBeUppercase, 0, "Test", nil, result{executeCommand: false}},

		{"cmdMustBeLowercase", MustBeLowercase, 0, "test", nil, result{executeCommand: true}},
		{"cmdMustBeLowercaseFail", MustBeLowercase, 0, "TEST", nil, result{executeCommand: false}},
		{"cmdMustBeLowercaseFailCapital", MustBeLowercase, 0, "Test", nil, result{executeCommand: false}},

		{"cmdMustHavePrefix", MustHavePrefix, 0, "testrakstar", "test", result{executeCommand: true}},
		{"cmdMustHavePrefixFail", MustHavePrefix, 0, "tastrakstar", "test", result{executeCommand: false}},
		{"cmdMustHavePrefixFailUppder", MustHavePrefix, 0, "TESTRAKSTAR", "test", result{executeCommand: false}},

		{"cmdMustHaveSufix", MustHaveSufix, 0, "rakstartest", "test", result{executeCommand: true}},
		{"cmdMustHaveSufixFail", MustHaveSufix, 0, "rakstartast", "test", result{executeCommand: false}},
		{"cmdMustHaveSufixFailUpper", MustHaveSufix, 0, "RAKSTARTEST", "test", result{executeCommand: false}},

		{"cmdMustCompileRegex", MustCompileRegex, 0, "support@rakstar.com", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, result{executeCommand: true}},
		{"cmdMustCompileRegexFail", MustCompileRegex, 0, "support", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, result{executeCommand: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := Builder()
			cmd.Conditionals().Index(tc.index)
			switch tc.typeTest {
			case MustBeUppercase:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeText().
					MustBeUppercase().
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
			case MustBeLowercase:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeText().
					MustBeLowercase().
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
			case MustHavePrefix:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeText().
					MustHavePrefix(tc.condValue.(string)).
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
			case MustHaveSufix:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeText().
					MustHaveSufix(tc.condValue.(string)).
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
			case MustCompileRegex:
				cmd.
					Command(tc.name).
					Handler(func(context *CommandContext) {}).
					Conditionals().
					Index(tc.index).
					TypeText().
					MustCompileRegex(tc.condValue.(string)).
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
			}
		})
	}
}

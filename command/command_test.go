package command

import (
	"fmt"
	"testing"
)

func TestRegisterCommand(t *testing.T) {
	commands = make(map[string]*Command)

	RegisterCommand(&Command{
		Name:    "test",
		Handler: func(context *CommandContext) {},
	})

	if len(commands) != 1 {
		t.Errorf("expected: 1; got %v", commands)
		return
	}
}

func TestRegisterCommandOverload(t *testing.T) {
	commands = make(map[string]*Command)

	for i := 0; i < 10000; i++ {
		if len(commands) != i {
			t.Errorf("expected: %v; got %v", i*2, commands)
			return
		}

		RegisterCommand(&Command{
			Name:    fmt.Sprintf("%dtest", i),
			Handler: func(context *CommandContext) {},
		})
	}
}

func TestRegisterCommandOverloadWithAliases(t *testing.T) {
	commands = make(map[string]*Command)

	for i := 0; i < 10000; i++ {
		if len(commands) != i*100+i {
			t.Errorf("expected: %v; got %v", i*100+i, commands)
			return
		}

		aliases := []string{}

		for l := 0; l < 100; l++ {
			aliases = append(aliases, fmt.Sprintf("%dalone%d", i, l))
		}

		RegisterCommand(&Command{
			Name:    fmt.Sprintf("%dtest", i),
			Aliases: aliases,
			Handler: func(context *CommandContext) {},
		})
	}
}

func TestSearchCommand(t *testing.T) {
	commands = make(map[string]*Command)

	var command *Command
	distance := -1

	command, distance = SearchCommand("test")

	if command != nil {
		t.Errorf("expected: nil; got %v", command)
		return
	}

	if distance != -1 {
		t.Errorf("expected: -1; got %d", distance)
		return
	}

	Builder().
		Command("test").
		Alias("exec").
		Handler(func(context *CommandContext) {}).
		Create()

	command = nil
	distance = -1

	command, distance = SearchCommand("test")

	if command == nil {
		t.Errorf("expected: defined; got %v", command)
		return
	}

	if distance != 0 {
		t.Errorf("expected: 0; got %d", distance)
		return
	}

	command = nil
	distance = -1

	command, distance = SearchCommand("exec")

	if command == nil {
		t.Errorf("expected: defined; got %v", command)
		return
	}

	if distance != 0 {
		t.Errorf("expected: 0; got %d", distance)
		return
	}

	command = nil
	distance = -1

	command, distance = SearchCommand("tast")

	if command == nil {
		t.Errorf("expected: defined; got %v", command)
		return
	}

	if distance != 1 {
		t.Errorf("expected: 1; got %d", distance)
		return
	}

	command, distance = SearchCommand("golang")

	if command != nil {
		t.Errorf("expected: nil; got %v", command)
		return
	}

	if distance < 2 {
		t.Errorf("expected: > 2; got %d", distance)
		return
	}
}

func TestSearchCommandOverload(t *testing.T) {
	commands = make(map[string]*Command)

	for i := 0; i < 10000; i++ {
		Builder().
			Command(fmt.Sprintf("%dtest", i)).
			Alias(fmt.Sprintf("%dexec", i)).
			Handler(func(context *CommandContext) {}).
			Create()

		command, distance := SearchCommand(fmt.Sprintf("%dtest", i))

		if command == nil {
			t.Errorf("expected: defined; got %v", command)
			return
		}

		if distance != 0 {
			t.Errorf("expected: 0; got %v", command)
			return
		}

		command, distance = SearchCommand(fmt.Sprintf("%dtest", i))

		if command == nil {
			t.Errorf("expected: defined; got %v", command)
			return
		}

		if distance != 0 {
			t.Errorf("expected: 0; got %v", command)
			return
		}
	}
}

func TestBuilderRegister(t *testing.T) {
	commands = make(map[string]*Command)
	Builder().
		Command("test").
		Handler(func(context *CommandContext) {}).
		Create()

	if len(commands) != 1 {
		t.Errorf("expected: 1; got %v", commands)
		return
	}
}

func TestBuilderCommandOverload(t *testing.T) {
	commands = make(map[string]*Command)

	for i := 0; i < 10000; i++ {
		if len(commands) != i {
			t.Errorf("expected: %v; got %v", i*2, commands)
			return
		}

		Builder().
			Command(fmt.Sprintf("%dtest", i)).
			Handler(func(context *CommandContext) {}).
			Create()
	}
}

func TestBuilderCommandOverloadWithAliases(t *testing.T) {
	commands = make(map[string]*Command)

	for i := 0; i < 10000; i++ {
		if len(commands) != i*100+i {
			t.Errorf("expected: %v; got %v", i*100+i, commands)
			return
		}

		cmdBuilder := Builder()
		cmdBuilder.
			Command(fmt.Sprintf("%dtest", i)).
			Handler(func(context *CommandContext) {})

		for l := 0; l < 100; l++ {
			cmdBuilder.Alias(fmt.Sprintf("%dalone%d", i, l))
		}

		cmdBuilder.Create()
	}
}

func TestConditionals(t*testing.T){
   
}

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliApp struct {
	commands Commands
	onStart  []CommandType
	canQuit  bool
}

func NewCliApp(commands Commands, onStart ...CommandType) *CliApp {
	return &CliApp{
		commands: commands,
		onStart:  onStart,
	}
}

func (app *CliApp) Run() error {
	for _, cmdtype := range app.onStart {
		cmd := app.commands[cmdtype]
		if err := app.executeCommand(cmd); err != nil {
			fmt.Println(err)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		if app.canQuit {
			return nil
		}
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		command, err := app.matchCommand(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// execute command
		if err := app.executeCommand(command); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (app *CliApp) matchCommand(input string) (*Command, error) {
	input = strings.TrimSpace(input)
	var command *Command
	for _, cmd := range app.commands {
		if cmd.Match(input) {
			command = cmd
			break
		}
	}
	if command == nil {
		return nil, fmt.Errorf("command for input %q not found", input)
	}
	matches := command.Regex.FindStringSubmatch(input)
	names := command.Regex.SubexpNames()
	if len(names) == 0 || len(matches[1:]) != len(names[1:]) {
		return nil, fmt.Errorf("command %q pattern must have %v named params", input, len(matches[1:]))
	}
	command.Params = make(map[string]string)
	for i := 1; i < len(names); i++ {
		command.Params[names[i]] = matches[i]
	}
	return command, nil
}

func (app *CliApp) executeCommand(cmd *Command) error {
	switch cmd.Type {
	case QuitCommand:
		app.canQuit = true
	default:
		if cmd.Controller == nil {
			return fmt.Errorf("execute command %v: nil controller", cmd)
		}
		err := cmd.Controller.Handle(cmd.Params)
		if err != nil {
			return fmt.Errorf("execute command %v: %v", cmd, err)
		}
	}
	return nil
}

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliApp struct {
	commands Commands
	canQuit  bool
}

func NewCliApp(commands Commands) *CliApp {
	return &CliApp{
		commands: commands,
	}
}

func (app *CliApp) Run() error {
	app.info()
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

func (app *CliApp) info() {
	fmt.Print(infoTemplate)
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
	case InfoCommand:
		app.info()
	default:
		return cmd.Controller.Handle(cmd.Params)
	}
	return nil
}

package cli

import (
	"regexp"

	"github.com/dmitruk-v/piggy-bank/internal/presentation/controllers"
)

type CommandType int

const (
	LoadBalance = iota
	ShowHelpCommand
	QuitCommand
	DepositCommand
	WithdrawCommand
	ShowBalanceCommand
	ShowOperationsCommand
	UndoCommand
)

type Command struct {
	Type       CommandType
	Regex      *regexp.Regexp
	Controller controllers.CliController
	Params     map[string]string
}

type Commands []*Command

func NewCommand(ctype CommandType, pattern string, ctrl controllers.CliController) *Command {
	return &Command{
		Type:       ctype,
		Regex:      regexp.MustCompile(pattern),
		Controller: ctrl,
	}
}

func (cmd *Command) Match(input string) bool {
	return cmd.Regex.MatchString(input)
}

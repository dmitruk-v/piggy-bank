package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
	"github.com/dmitruk-v/piggy-bank/internal/usecase"
)

type CliShowOperationsController struct {
	showOpsUcase usecase.ShowOperationsUseCaseInput
}

func NewCliShowOperationsController(showOpsUcase usecase.ShowOperationsUseCaseInput) *CliShowOperationsController {
	return &CliShowOperationsController{
		showOpsUcase: showOpsUcase,
	}
}

func (ctrl *CliShowOperationsController) Handle(req CliRequest) error {
	res := ctrl.showOpsUcase.Execute()
	if res.Err != nil {
		return res.Err
	}
	b := strings.Builder{}
	b.WriteString("\n")
	for i, op := range res.Operations {
		b.WriteString(" ")
		switch op.Optype {
		case domain.DepositOperation:
			b.WriteString(fmt.Sprintf("%v. \x1b[32m", i))
			b.WriteString("+")
		case domain.WithdrawOperation:
			b.WriteString(fmt.Sprintf("%v. \x1b[42m", i))
			b.WriteString("-")
		}
		b.WriteString(" ")
		b.WriteString(strings.ToUpper(string(op.Currency)))
		b.WriteString(" ")
		b.WriteString(fmt.Sprintf("%v", op.Amount))
		b.WriteString(" ")
		optime := time.UnixMilli(op.ProvidedAt).Format("02/01/2006 15:04:05")
		b.WriteString(fmt.Sprintf("%v", optime))
		b.WriteString("\x1b[0m")
		b.WriteString("\n")
	}
	b.WriteString("\n")
	fmt.Print(b.String())
	return nil
}

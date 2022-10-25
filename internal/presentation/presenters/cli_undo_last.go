package presenters

import (
	"fmt"
	"io"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
)

type CliUndoLastPresenter struct {
	writer io.Writer
}

func NewCliUndoLastPresenter(w io.Writer) *CliUndoLastPresenter {
	return &CliUndoLastPresenter{
		writer: w,
	}
}

func (pr *CliUndoLastPresenter) Present(res usecase.UndoLastResponse) {
	if res.Error != nil {
		fmt.Fprintf(pr.writer, "Error: %v\n", res.Error)
		return
	}
	op := res.Operation
	sign := ""
	switch op.Optype {
	case entity.DepositOperation:
		sign = "+"
	case entity.WithdrawOperation:
		sign = "-"
	}
	when := time.Unix(op.ProvidedAt, 0).Format(TimeFormat)
	fmt.Fprintf(pr.writer, "Undone operation [%v %v %v] provided at %v\n", sign, op.Amount, op.Currency, when)
}

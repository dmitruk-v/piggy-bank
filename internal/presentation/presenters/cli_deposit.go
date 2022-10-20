package presenters

import (
	"fmt"
	"io"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
)

type CliDepositPresenter struct {
	writer io.Writer
}

func NewCliDepositPresenter(w io.Writer) *CliDepositPresenter {
	return &CliDepositPresenter{
		writer: w,
	}
}

func (pr *CliDepositPresenter) Present(res usecase.DepositResponse) {
	if res.Error != nil {
		fmt.Fprintf(pr.writer, "Error: %v\n", res.Error)
		return
	}
	op := res.Operation
	switch op.Optype {
	case entity.DepositOperation:
		fmt.Fprint(pr.writer, "+ ")
	case entity.WithdrawOperation:
		fmt.Fprint(pr.writer, "- ")
	}
	fmt.Fprintf(pr.writer, "%v %v\n", op.Amount, op.Currency)
}

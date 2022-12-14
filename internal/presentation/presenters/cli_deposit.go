package presenters

import (
	"fmt"
	"io"
	"time"

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
	when := time.Unix(op.ProvidedAt, 0).Format(TimeFormat)
	fmt.Fprintf(pr.writer, "[+ %v %v] at %v\n", op.Currency, op.Amount, when)
}

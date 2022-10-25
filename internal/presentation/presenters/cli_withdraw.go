package presenters

import (
	"fmt"
	"io"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
)

type CliWithdrawPresenter struct {
	writer io.Writer
}

func NewCliWithdrawPresenter(w io.Writer) *CliWithdrawPresenter {
	return &CliWithdrawPresenter{
		writer: w,
	}
}

func (pr *CliWithdrawPresenter) Present(res usecase.WithdrawResponse) {
	if res.Error != nil {
		fmt.Fprintf(pr.writer, "Error: %v\n", res.Error)
		return
	}
	op := res.Operation
	when := time.Unix(op.ProvidedAt, 0).Format(TimeFormat)
	fmt.Fprintf(pr.writer, "[- %v %v] at %v\n", op.Currency, op.Amount, when)
}

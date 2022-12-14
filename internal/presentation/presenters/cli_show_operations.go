package presenters

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
)

type ShowOperationUseCaseOutput interface {
	Present(res usecase.ShowOperationsResponse)
}

type CliShowOperationPresenter struct {
	writer io.Writer
}

func NewCliShowOperationsPresenter(w io.Writer) *CliShowOperationPresenter {
	return &CliShowOperationPresenter{
		writer: w,
	}
}

func (pr *CliShowOperationPresenter) Present(res usecase.ShowOperationsResponse) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape)
	fmt.Fprintln(tw)
	fmt.Fprintln(tw, "#\tOperation\tCurrency\tAmount\tProvidedAt")
	fmt.Fprintln(tw, "---\t---------\t--------\t------\t----------")
	sign := "nosign"
	for i, op := range res.Operations {
		switch op.Optype {
		case entity.DepositOperation:
			sign = "+"
		case entity.WithdrawOperation:
			sign = "-"
		}
		optime := time.Unix(op.ProvidedAt, 0).Format("02/01/2006 15:04:05")
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\t%v\n", i+1, sign, op.Currency, op.Amount, optime)
	}
	fmt.Fprintln(tw)
	tw.Flush()
}

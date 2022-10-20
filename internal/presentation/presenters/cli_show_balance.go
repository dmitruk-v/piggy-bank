package presenters

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
)

type CliShowBalancePresenter struct {
	writer io.Writer
}

func NewCliShowBalancePresenter(w io.Writer) *CliShowBalancePresenter {
	return &CliShowBalancePresenter{
		writer: w,
	}
}

func (pr *CliShowBalancePresenter) Present(res usecase.ShowBalanceResponse) error {
	tw := tabwriter.NewWriter(pr.writer, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw)
	fmt.Fprintln(tw, "#\tCurrency\tAmount")
	fmt.Fprintln(tw, "---\t--------\t------")
	for i, item := range res.Balance.List() {
		fmt.Fprintf(tw, "%v\t%v\t%v\n", i, item.Curr, item.Amount)
	}
	fmt.Fprintln(tw)
	return tw.Flush()
}

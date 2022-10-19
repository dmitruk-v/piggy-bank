package presenters

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

type CliShowHelpPresenter struct {
	writer io.Writer
}

func NewCliShowHelpPresenter(w io.Writer) *CliShowHelpPresenter {
	return &CliShowHelpPresenter{
		writer: w,
	}
}

func (pr *CliShowHelpPresenter) Present() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw)
	fmt.Fprintln(tw, "#\tCommand\tArguments\tInfo")
	fmt.Fprintln(tw, "---\t-------\t---------\t----")
	fmt.Fprintln(tw, "1\thelp\t\tInformation about commands usage")
	fmt.Fprintln(tw, "2\tquit\t\tQuit from app")
	fmt.Fprintln(tw, "3\tdeposit\tCURRENCY AMOUNT\tAdd amount of currency to balance")
	fmt.Fprintln(tw, "4\twithdraw\tCURRENCY AMOUNT\tSubstract amount of currency from balance")
	fmt.Fprintln(tw, "5\tbalance\t\tInfo about balance")
	fmt.Fprintln(tw, "6\toperations\t\tInfo about operations")
	fmt.Fprintln(tw, "7\tundo\t\tUndo last operation")
	fmt.Fprintln(tw)
	tw.Flush()
}

package controllers

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
	"github.com/dmitruk-v/piggy-bank/internal/usecase"
)

type CliWithdrawController struct {
	withdrawUcase usecase.WithdrawUseCaseInput
}

func NewCliWithdrawController(withdrawUcase usecase.WithdrawUseCaseInput) *CliWithdrawController {
	return &CliWithdrawController{
		withdrawUcase: withdrawUcase,
	}
}

func (ctrl *CliWithdrawController) Handle(req CliRequest) error {
	curr := req["currency"]
	amt := req["amount"]
	// TODO: validate request params

	fmt.Println(curr, amt)
	res := ctrl.withdrawUcase.Execute(usecase.WithdrawRequest{
		Currency: domain.EUR,
		Amount:   123.45,
	})
	_ = res
	return nil
}
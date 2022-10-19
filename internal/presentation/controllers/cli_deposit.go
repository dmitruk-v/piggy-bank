package controllers

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
	"github.com/dmitruk-v/piggy-bank/internal/usecase"
)

type CliDepositController struct {
	depositUcase usecase.DepositUseCaseInput
}

func NewCliDepositController(depositUcase usecase.DepositUseCaseInput) *CliDepositController {
	return &CliDepositController{
		depositUcase: depositUcase,
	}
}

func (ctrl *CliDepositController) Handle(req CliRequest) error {
	curr := req["currency"]
	amt := req["amount"]
	// TODO: validate request params

	fmt.Println(curr, amt)
	res := ctrl.depositUcase.Execute(usecase.DepositRequest{
		Currency: domain.EUR,
		Amount:   123.45,
	})
	_ = res
	return nil
}

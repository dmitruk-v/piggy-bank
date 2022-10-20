package controllers

import (
	"strconv"
	"strings"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
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
	curr := entity.Currency(strings.ToUpper(req["currency"]))
	amtVal, err := strconv.ParseFloat(req["amount"], 64)
	if err != nil {
		return err
	}
	ctrl.depositUcase.Execute(usecase.DepositRequest{
		Currency: curr,
		Amount:   amtVal,
	})
	return nil
}

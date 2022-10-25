package controllers

import (
	"fmt"
	"strconv"

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
	depositReq, err := ctrl.validateRequest(req)
	if err != nil {
		return err
	}
	ctrl.depositUcase.Execute(depositReq)
	return nil
}

func (ctrl *CliDepositController) validateRequest(req CliRequest) (usecase.DepositRequest, error) {
	fmtError := func(err error) error {
		return fmt.Errorf("validate cli deposit request: %v", err)
	}
	var outReq usecase.DepositRequest
	curr, err := entity.CurrencyFromString(req["currency"])
	if err != nil {
		return outReq, fmtError(err)
	}
	amt, err := strconv.ParseFloat(req["amount"], 64)
	if err != nil {
		return outReq, fmtError(err)
	}
	outReq.Currency = curr
	outReq.Amount = amt
	return outReq, nil
}

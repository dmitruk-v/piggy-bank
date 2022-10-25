package controllers

import (
	"fmt"
	"strconv"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
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
	withdrawReq, err := ctrl.validateRequest(req)
	if err != nil {
		return err
	}
	ctrl.withdrawUcase.Execute(withdrawReq)
	return nil
}

func (ctrl *CliWithdrawController) validateRequest(req CliRequest) (usecase.WithdrawRequest, error) {
	fmtError := func(err error) error {
		return fmt.Errorf("validate cli withdraw request: %v", err)
	}
	var outReq usecase.WithdrawRequest
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

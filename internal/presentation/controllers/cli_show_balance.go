package controllers

import "github.com/dmitruk-v/piggy-bank/internal/usecase"

type CliShowBalanceController struct {
	showBalanceUcase usecase.ShowBalanceUseCaseInput
}

func NewCliShowBalanceController(showBalanceUcase usecase.ShowBalanceUseCaseInput) *CliShowBalanceController {
	return &CliShowBalanceController{
		showBalanceUcase: showBalanceUcase,
	}
}

func (ctrl *CliShowBalanceController) Execute() error {
	return nil
}

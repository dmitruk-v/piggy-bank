package controllers

import "github.com/dmitruk-v/piggy-bank/internal/domain/usecase"

type CliShowBalanceController struct {
	showBalanceUcase usecase.ShowBalanceUseCaseInput
}

func NewCliShowBalanceController(showBalanceUcase usecase.ShowBalanceUseCaseInput) *CliShowBalanceController {
	return &CliShowBalanceController{
		showBalanceUcase: showBalanceUcase,
	}
}

func (ctrl *CliShowBalanceController) Handle(req CliRequest) error {
	ctrl.showBalanceUcase.Execute()
	return nil
}

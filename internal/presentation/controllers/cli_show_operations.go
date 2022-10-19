package controllers

import "github.com/dmitruk-v/piggy-bank/internal/domain/usecase"

type CliShowOperationsController struct {
	showOpsUcase usecase.ShowOperationsUseCaseInput
}

func NewCliShowOperationsController(showOpsUcase usecase.ShowOperationsUseCaseInput) *CliShowOperationsController {
	return &CliShowOperationsController{
		showOpsUcase: showOpsUcase,
	}
}

func (ctrl *CliShowOperationsController) Handle(req CliRequest) error {
	ctrl.showOpsUcase.Execute()
	return nil
}

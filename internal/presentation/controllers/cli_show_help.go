package controllers

import "github.com/dmitruk-v/piggy-bank/internal/domain/usecase"

type CliShowHelpController struct {
	showHelpUcase usecase.ShowHelpUseCaseInput
}

func NewCliShowHelpController(showHelpUcase usecase.ShowHelpUseCaseInput) *CliShowHelpController {
	return &CliShowHelpController{
		showHelpUcase: showHelpUcase,
	}
}

func (ctrl *CliShowHelpController) Handle(req CliRequest) error {
	ctrl.showHelpUcase.Execute()
	return nil
}

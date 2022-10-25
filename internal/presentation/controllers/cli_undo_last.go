package controllers

import "github.com/dmitruk-v/piggy-bank/internal/domain/usecase"

type CliUndoLastController struct {
	undoLastUcase usecase.UndoLastUseCaseInput
}

func NewCliUndoLastController(undoLastUcase usecase.UndoLastUseCaseInput) *CliUndoLastController {
	return &CliUndoLastController{
		undoLastUcase: undoLastUcase,
	}
}

func (ctrl *CliUndoLastController) Handle(req CliRequest) error {
	ctrl.undoLastUcase.Execute()
	return nil
}

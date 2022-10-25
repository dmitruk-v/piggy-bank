package controllers

import "github.com/dmitruk-v/piggy-bank/internal/domain/usecase"

type CliLoadBalanceController struct {
	loadBalanceUcase usecase.LoadBalanceUseCaseInput
}

func NewCliLoadBalanceController(loadBalanceUcase usecase.LoadBalanceUseCaseInput) *CliLoadBalanceController {
	return &CliLoadBalanceController{
		loadBalanceUcase: loadBalanceUcase,
	}
}

func (ctrl *CliLoadBalanceController) Handle(req CliRequest) error {
	return ctrl.loadBalanceUcase.Execute()
}

package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type ShowOperationsResponse struct {
	Err        error
	Operations []*entity.CurrencyOperation
}

type ShowOperationsUseCaseInput interface {
	Execute()
}

type ShowOperationsUseCaseOutput interface {
	Present(res ShowOperationsResponse)
}

type ShowOperationsUseCase struct {
	opStorage entity.OperationStorage
	output    ShowOperationsUseCaseOutput
}

func NewShowOperationsUseCase(opStorage entity.OperationStorage, output ShowOperationsUseCaseOutput) *ShowOperationsUseCase {
	return &ShowOperationsUseCase{
		opStorage: opStorage,
		output:    output,
	}
}

func (ucase *ShowOperationsUseCase) Execute() {
	var res ShowOperationsResponse
	ops, err := ucase.opStorage.GetAll()
	if err != nil {
		res.Err = fmt.Errorf("execute show operations use case: %v", err)
		ucase.output.Present(res)
		return
	}
	res.Operations = ops
	ucase.output.Present(res)
}

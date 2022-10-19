package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type ShowOperationsResponse struct {
	Err        error
	Operations []*domain.CurrencyOperation
}

type ShowOperationsUseCaseInput interface {
	Execute() ShowOperationsResponse
}

type ShowOperationsUseCase struct {
	opStorage domain.OperationStorage
}

func NewShowOperationsUseCase(opStorage domain.OperationStorage) *ShowOperationsUseCase {
	return &ShowOperationsUseCase{
		opStorage: opStorage,
	}
}

func (ucase *ShowOperationsUseCase) Execute() ShowOperationsResponse {
	var res ShowOperationsResponse
	ops, err := ucase.opStorage.GetAll()
	if err != nil {
		res.Err = fmt.Errorf("execute info operations use case: %v", err)
		return res
	}
	res.Operations = ops
	return res
}

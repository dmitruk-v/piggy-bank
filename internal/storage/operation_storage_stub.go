package storage

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type StubOperationStorage struct {
	operations []*domain.CurrencyOperation
}

func NewStubOperationStorage() *StubOperationStorage {
	return &StubOperationStorage{
		operations: make([]*domain.CurrencyOperation, 0),
	}
}

func (stg *StubOperationStorage) List() ([]*domain.CurrencyOperation, error) {
	return stg.operations, nil
}

func (stg *StubOperationStorage) Save(op *domain.CurrencyOperation) error {
	stg.operations = append(stg.operations, op)
	return nil
}

func (stg *StubOperationStorage) DeleteLast() (*domain.CurrencyOperation, error) {
	if len(stg.operations) == 0 {
		return nil, fmt.Errorf("there are no operations yet")
	}
	op := stg.operations[len(stg.operations)-1]
	stg.operations = stg.operations[:len(stg.operations)-1]
	return op, nil
}

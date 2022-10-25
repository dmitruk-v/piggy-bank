package storage

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type StubOperationStorage struct {
	operations []*entity.CurrencyOperation
}

func NewStubOperationStorage(initial []*entity.CurrencyOperation) *StubOperationStorage {
	return &StubOperationStorage{
		operations: initial,
	}
}

func (stg *StubOperationStorage) GetAll() ([]*entity.CurrencyOperation, error) {
	return stg.operations, nil
}

func (stg *StubOperationStorage) GetLatest() (*entity.CurrencyOperation, error) {
	if len(stg.operations) == 0 {
		return nil, entity.ErrNoOperations
	}
	return stg.operations[len(stg.operations)-1], nil
}

func (stg *StubOperationStorage) Save(op *entity.CurrencyOperation) error {
	stg.operations = append(stg.operations, op)
	return nil
}

func (stg *StubOperationStorage) DeleteLatest() (*entity.CurrencyOperation, error) {
	if len(stg.operations) == 0 {
		return nil, fmt.Errorf("there are no operations yet")
	}
	op := stg.operations[len(stg.operations)-1]
	stg.operations = stg.operations[:len(stg.operations)-1]
	return op, nil
}

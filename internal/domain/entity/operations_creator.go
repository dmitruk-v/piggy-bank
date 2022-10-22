package entity

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/common"
)

type OperationCreator interface {
	Create(optype OperationType, currency Currency, amount float64, providedAt int64, prevHash []byte) (*CurrencyOperation, error)
}

type OperationsCreatorImpl struct {
	bcService common.BlockchainService
}

func NewOperationsCreatorImpl(bcService common.BlockchainService) *OperationsCreatorImpl {
	return &OperationsCreatorImpl{
		bcService: bcService,
	}
}

func (ctor *OperationsCreatorImpl) Create(optype OperationType, currency Currency, amount float64, providedAt int64, prevHash []byte) (*CurrencyOperation, error) {
	opstr := fmt.Sprintf("%v%v%v%v%v", optype, currency, amount, providedAt, prevHash)
	hash, err := ctor.bcService.Hash([]byte(opstr))
	if err != nil {
		return nil, fmt.Errorf("create operation: %v", err)
	}
	op := NewCurrencyOperation(optype, currency, amount, providedAt, hash, prevHash)
	return op, nil
}

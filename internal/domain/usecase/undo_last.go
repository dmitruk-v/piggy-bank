package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type UndoLastUseCase struct {
	balance   entity.Balance
	opStorage entity.OperationStorage
}

func NewUndoLastUseCase(balance entity.Balance, opStorage entity.OperationStorage) *UndoLastUseCase {
	return &UndoLastUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *UndoLastUseCase) Execute() error {
	fmtError := func(err error) error {
		return fmt.Errorf("execute undo-last operation: %v", err)
	}
	op, err := ucase.opStorage.DeleteLatest()
	if err != nil {
		return fmtError(err)
	}
	switch op.Optype {
	case entity.DepositOperation:
		if err := ucase.balance.Sub(op.Currency, op.Amount); err != nil {
			return fmtError(err)
		}
	case entity.WithdrawOperation:
		if err := ucase.balance.Add(op.Currency, op.Amount); err != nil {
			return fmtError(err)
		}
	}
	return nil
}

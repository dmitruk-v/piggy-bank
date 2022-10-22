package usecase

import (
	"errors"
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
	ops, err := ucase.opStorage.PopLatest(1)
	if err != nil {
		return fmtError(err)
	}
	if len(ops) == 0 {
		return fmtError(errors.New("no operations found"))
	}
	for _, op := range ops {
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
	}
	return nil
}

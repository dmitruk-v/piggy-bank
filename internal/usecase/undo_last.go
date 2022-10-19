package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type UndoLastUseCase struct {
	balance   *domain.Balance
	opStorage domain.OperationStorage
}

func NewUndoLastUseCase(balance *domain.Balance, opStorage domain.OperationStorage) *UndoLastUseCase {
	return &UndoLastUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *UndoLastUseCase) Execute() error {
	op, err := ucase.opStorage.DeleteLast()
	if err != nil {
		return fmtError(err)
	}
	switch op.Optype {
	case domain.DepositOperation:
		if err := ucase.balance.Sub(op.Currency, op.Amount); err != nil {
			return fmtError(err)
		}
	case domain.WithdrawOperation:
		if err := ucase.balance.Add(op.Currency, op.Amount); err != nil {
			return fmtError(err)
		}
	}
	return nil
}

func fmtError(err error) error {
	return fmt.Errorf("execute undo-last operation: %v", err)
}

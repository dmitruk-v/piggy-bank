package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type UndoLastUseCaseInput interface {
	Execute()
}

type UndoLastResponse struct {
	Operation *entity.CurrencyOperation
	Error     error
}

type UndoLastUseCaseOutput interface {
	Present(res UndoLastResponse)
}

type UndoLastUseCase struct {
	balance   entity.Balance
	opStorage entity.OperationStorage
	output    UndoLastUseCaseOutput
}

func NewUndoLastUseCase(balance entity.Balance, opStorage entity.OperationStorage, output UndoLastUseCaseOutput) *UndoLastUseCase {
	return &UndoLastUseCase{
		balance:   balance,
		opStorage: opStorage,
		output:    output,
	}
}

func (ucase *UndoLastUseCase) Execute() {
	makeError := func(err error) error {
		return fmt.Errorf("execute undo-last operation: %v", err)
	}
	var res UndoLastResponse
	defer func() {
		ucase.output.Present(res)
	}()
	op, err := ucase.opStorage.GetLatest()
	if err != nil {
		res.Error = makeError(err)
		return
	}
	switch op.Optype {
	case entity.DepositOperation:
		if err := ucase.balance.Sub(op.Currency, op.Amount); err != nil {
			res.Error = makeError(err)
			return
		}
	case entity.WithdrawOperation:
		if err := ucase.balance.Add(op.Currency, op.Amount); err != nil {
			res.Error = makeError(err)
			return
		}
	}
	op, err = ucase.opStorage.DeleteLatest()
	if err != nil {
		res.Error = makeError(err)
		return
	}
	res.Operation = op
}

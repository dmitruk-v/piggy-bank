package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type DepositRequest struct {
	Currency entity.Currency
	Amount   float64
}

type DepositUseCaseInput interface {
	Execute(req DepositRequest)
}

type DepositUseCaseOutput interface {
	Present(res DepositResponse)
}

type DepositResponse struct {
	Operation *entity.CurrencyOperation
	Error     error
}

type DepositUseCase struct {
	balance   entity.Balance
	opStorage entity.OperationStorage
	opCreator entity.OperationCreator
	output    DepositUseCaseOutput
}

func NewDepositUseCase(balance entity.Balance, opStorage entity.OperationStorage, opCreator entity.OperationCreator, output DepositUseCaseOutput) *DepositUseCase {
	return &DepositUseCase{
		balance:   balance,
		opStorage: opStorage,
		opCreator: opCreator,
		output:    output,
	}
}

func (ucase *DepositUseCase) Execute(req DepositRequest) {
	fmtError := func(err error) error {
		return fmt.Errorf("execute deposit operation: %v", err)
	}
	var res DepositResponse
	defer func() {
		ucase.output.Present(res)
	}()
	var prevHash []byte
	lop, err := ucase.opStorage.GetLatest()
	if err != nil {
		if !errors.Is(err, entity.ErrNoOperations) {
			res.Error = fmtError(err)
			return
		}
	}
	if lop != nil {
		prevHash = lop.Hash
	} else {
		prevHash = nil
	}
	op, err := ucase.opCreator.Create(entity.DepositOperation, req.Currency, req.Amount, time.Now().Unix(), prevHash)
	if err != nil {
		res.Error = fmtError(err)
		return
	}
	res.Operation = op
	if err := ucase.opStorage.Save(op); err != nil {
		res.Error = fmtError(err)
		return
	}
	if err := ucase.balance.Add(req.Currency, req.Amount); err != nil {
		res.Error = fmtError(err)
		return
	}
}

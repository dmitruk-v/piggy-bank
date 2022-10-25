package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type WithdrawRequest struct {
	Currency entity.Currency
	Amount   float64
}

type WithdrawUseCaseInput interface {
	Execute(req WithdrawRequest)
}

type WithdrawResponse struct {
	Operation *entity.CurrencyOperation
	Error     error
}

type WithdrawUseCaseOutput interface {
	Present(res WithdrawResponse)
}

type WithdrawUseCase struct {
	balance   entity.Balance
	opStorage entity.OperationStorage
	opCreator entity.OperationCreator
	output    WithdrawUseCaseOutput
}

func NewWithdrawUseCase(balance entity.Balance, opStorage entity.OperationStorage, opCreator entity.OperationCreator, output WithdrawUseCaseOutput) *WithdrawUseCase {
	return &WithdrawUseCase{
		balance:   balance,
		opStorage: opStorage,
		opCreator: opCreator,
		output:    output,
	}
}

func (ucase *WithdrawUseCase) Execute(req WithdrawRequest) {
	fmtError := func(err error) error {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	var res WithdrawResponse
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
	op, err := ucase.opCreator.Create(entity.WithdrawOperation, req.Currency, req.Amount, time.Now().Unix(), prevHash)
	if err != nil {
		res.Error = fmtError(err)
		return
	}
	res.Operation = op
	if err := ucase.opStorage.Save(op); err != nil {
		res.Error = fmtError(err)
		return
	}
	if err := ucase.balance.Sub(req.Currency, req.Amount); err != nil {
		res.Error = fmtError(err)
		return
	}
}

package usecase

import (
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/common"
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
	bcService common.BlockchainService
	output    DepositUseCaseOutput
}

func NewDepositUseCase(balance entity.Balance, opStorage entity.OperationStorage, bcService common.BlockchainService, output DepositUseCaseOutput) *DepositUseCase {
	return &DepositUseCase{
		balance:   balance,
		opStorage: opStorage,
		bcService: bcService,
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
	lops, err := ucase.opStorage.GetLatest(1)
	if err != nil {
		res.Error = fmtError(err)
		return
	}
	var prevHash []byte
	if len(lops) > 0 {
		prevHash = lops[len(lops)-1].Hash
	}
	hash, err := ucase.bcService.Hash()
	if err != nil {
		res.Error = fmtError(err)
		return
	}
	//
	// TODO: make generated hash to depend on operation details
	//
	// if !ucase.balance.HasCurrency(req.Currency) {

	// }
	op := entity.NewCurrencyOperation(entity.DepositOperation, req.Currency, req.Amount, time.Now().Unix(), hash, prevHash)
	if err := ucase.opStorage.Save(op); err != nil {
		res.Error = fmtError(err)
		return
	}
	if err := ucase.balance.Add(req.Currency, req.Amount); err != nil {
		res.Error = fmtError(err)
		return
	}
	res.Operation = op
}

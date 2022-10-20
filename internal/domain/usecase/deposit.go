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
	Execute(req DepositRequest) DepositResponse
}

type DepositResponse struct {
	Err     error
	Message string
}

type DepositUseCase struct {
	balance   *entity.Balance
	opStorage entity.OperationStorage
	bcService common.BlockchainService
}

func NewDepositUseCase(balance *entity.Balance, opStorage entity.OperationStorage, bcService common.BlockchainService) *DepositUseCase {
	return &DepositUseCase{
		balance:   balance,
		opStorage: opStorage,
		bcService: bcService,
	}
}

func (ucase *DepositUseCase) Execute(req DepositRequest) DepositResponse {
	var (
		hash     []byte
		prevHash []byte
		res      DepositResponse
	)
	lops, err := ucase.opStorage.GetLatest(1)
	if err != nil {
		res.Err = fmt.Errorf("execute deposit operation: %v", err)
		return res
	}
	hash, err = ucase.bcService.Hash()
	if err != nil {
		res.Err = fmt.Errorf("execute deposit operation: %v", err)
		return res
	}
	if len(lops) > 0 {
		prevHash = lops[len(lops)-1].Hash
	}
	//
	// TODO: make generated hash to depend on operation details
	//
	op := entity.NewCurrencyOperation(entity.DepositOperation, req.Currency, req.Amount, time.Now().Unix(), hash, prevHash)
	if err := ucase.opStorage.Save(op); err != nil {
		res.Err = fmt.Errorf("execute deposit operation: %v", err)
		return res
	}
	if err := ucase.balance.Add(req.Currency, req.Amount); err != nil {
		res.Err = fmt.Errorf("execute deposit operation: %v", err)
		return res
	}
	res.Message = "Operation successfuly provided"
	return res
}

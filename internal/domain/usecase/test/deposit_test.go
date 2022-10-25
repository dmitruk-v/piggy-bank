package usecase

import (
	"bytes"
	"testing"

	"github.com/dmitruk-v/piggy-bank/internal/common"
	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
	"github.com/dmitruk-v/piggy-bank/internal/presentation/presenters"
	"github.com/dmitruk-v/piggy-bank/internal/storage"
)

func TestDepositUseCase(t *testing.T) {
	bcService := common.NewBlockchainServiceImpl()
	balance := entity.NewBalanceImpl(entity.EUR, entity.UAH)
	opStorage := storage.NewStubOperationStorage(nil)
	opCreator := entity.NewOperationsCreatorImpl(bcService)
	writer := bytes.NewBufferString("")
	presenter := presenters.NewCliDepositPresenter(writer)
	ucase := usecase.NewDepositUseCase(balance, opStorage, opCreator, presenter)
	req := usecase.DepositRequest{
		Currency: entity.EUR,
		Amount:   12500,
	}
	ucase.Execute(req)
	amount := balance.Amount(entity.EUR)
	if amount != req.Amount {
		t.Errorf("for currency %v, got amount: %v, want: %v", req.Currency, amount, req.Amount)
	}
}

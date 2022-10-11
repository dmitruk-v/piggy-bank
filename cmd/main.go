package main

import (
	"fmt"
	"log"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
	"github.com/dmitruk-v/piggy-bank/internal/storage"
	"github.com/dmitruk-v/piggy-bank/internal/usecase"
)

func main() {
	balance := domain.NewBalance(domain.USD, domain.EUR, domain.RUB, domain.UAH)

	opStorage := storage.NewFileOperationStorage("data.txt")
	depositUcase := usecase.NewDepositUseCase(balance, opStorage)
	withdrawUcase := usecase.NewWithdrawUseCase(balance, opStorage)
	// undoLastUcase := usecase.NewUndoLastUseCase(balance, opStorage)

	if err := depositUcase.Execute(domain.EUR, 125.0); err != nil {
		log.Fatal(err)
	}

	if err := withdrawUcase.Execute(domain.EUR, 50.0); err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	ops, err := opStorage.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, op := range ops {
		fmt.Printf("%#v\n", op)
	}
}

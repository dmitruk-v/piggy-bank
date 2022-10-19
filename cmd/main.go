package main

import (
	"fmt"
	"log"

	"github.com/dmitruk-v/piggy-bank/cmd/cli"
	"github.com/dmitruk-v/piggy-bank/internal/common"
	"github.com/dmitruk-v/piggy-bank/internal/domain"
	"github.com/dmitruk-v/piggy-bank/internal/presentation/controllers"

	"github.com/dmitruk-v/piggy-bank/internal/storage"
	"github.com/dmitruk-v/piggy-bank/internal/usecase"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	opStorage := storage.NewFileOperationStorage("operations.data")

	currencies := []domain.Currency{domain.USD, domain.EUR, domain.RUB, domain.UAH}
	balance := domain.NewBalance(currencies)

	loadBalanceUcase := usecase.NewLoadBalanceUseCase(balance, opStorage)
	if err := loadBalanceUcase.Execute(); err != nil {
		log.Fatal(err)
	}

	blockChainService := common.NewBlockchainServiceImpl()

	depositUcase := usecase.NewDepositUseCase(balance, opStorage, blockChainService)
	depositController := controllers.NewCliDepositController(depositUcase)
	withdrawUcase := usecase.NewWithdrawUseCase(balance, opStorage)
	withdrawController := controllers.NewCliWithdrawController(withdrawUcase)

	showOpsUcase := usecase.NewShowOperationsUseCase(opStorage)
	showOpsController := controllers.NewCliShowOperationsController(showOpsUcase)

	fmt.Println(balance)

	app := cli.NewCliApp(cli.Commands{
		cli.NewCommand(cli.InfoCommand, `^info$`, nil),
		cli.NewCommand(cli.QuitCommand, `^quit$`, nil),
		cli.NewCommand(cli.DepositCommand, `^deposit (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, depositController),
		cli.NewCommand(cli.WithdrawCommand, `^withdraw (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, withdrawController),
		cli.NewCommand(cli.ShowBalanceCommand, `^balance$`, nil),
		cli.NewCommand(cli.ShowOperationsCommand, `^operations$`, showOpsController),
		cli.NewCommand(cli.UndoCommand, `^undo$`, nil),
	})

	return app.Run()
}

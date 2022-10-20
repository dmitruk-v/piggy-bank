package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dmitruk-v/piggy-bank/cmd/cli"
	"github.com/dmitruk-v/piggy-bank/internal/common"
	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/domain/usecase"
	"github.com/dmitruk-v/piggy-bank/internal/presentation/controllers"
	"github.com/dmitruk-v/piggy-bank/internal/presentation/presenters"
	"github.com/dmitruk-v/piggy-bank/internal/storage"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	opStorage := storage.NewFileOperationStorage("operations.data")

	currencies := []entity.Currency{entity.USD, entity.EUR, entity.RUB, entity.UAH}
	balance := entity.NewBalance(currencies)

	loadBalanceUcase := usecase.NewLoadBalanceUseCase(balance, opStorage)
	if err := loadBalanceUcase.Execute(); err != nil {
		log.Fatal(err)
	}

	blockChainService := common.NewBlockchainServiceImpl()

	showHelpPresenter := presenters.NewCliShowHelpPresenter(os.Stdout)
	showHelpUcase := usecase.NewShowHelpUseCase(showHelpPresenter)
	showHelpController := controllers.NewCliShowHelpController(showHelpUcase)

	depositUcase := usecase.NewDepositUseCase(balance, opStorage, blockChainService)
	depositController := controllers.NewCliDepositController(depositUcase)

	withdrawUcase := usecase.NewWithdrawUseCase(balance, opStorage)
	withdrawController := controllers.NewCliWithdrawController(withdrawUcase)

	showBalancePresenter := presenters.NewCliShowBalancePresenter(os.Stdout)
	showBalanceUcase := usecase.NewShowBalanceUseCase(balance, showBalancePresenter)
	showBalanceController := controllers.NewCliShowBalanceController(showBalanceUcase)

	showOpsPresenter := presenters.NewCliShowOperationsPresenter(os.Stdout)
	showOpsUcase := usecase.NewShowOperationsUseCase(opStorage, showOpsPresenter)
	showOpsController := controllers.NewCliShowOperationsController(showOpsUcase)

	fmt.Println(balance)

	commands := cli.Commands{
		cli.NewCommand(cli.ShowHelpCommand, `^help$`, showHelpController),
		cli.NewCommand(cli.QuitCommand, `^quit$`, nil),
		cli.NewCommand(cli.DepositCommand, `^deposit (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, depositController),
		cli.NewCommand(cli.WithdrawCommand, `^withdraw (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, withdrawController),
		cli.NewCommand(cli.ShowBalanceCommand, `^balance$`, showBalanceController),
		cli.NewCommand(cli.ShowOperationsCommand, `^operations|ops$`, showOpsController),
		cli.NewCommand(cli.UndoCommand, `^undo$`, nil),
	}
	app := cli.NewCliApp(commands, commands[0])

	return app.Run()
}

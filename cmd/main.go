package main

import (
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
	balance := entity.NewBalanceImpl([]entity.Currency{
		entity.USD, entity.EUR, entity.RUB, entity.UAH,
	})
	blockChainService := common.NewBlockchainServiceImpl()
	opCreator := entity.NewOperationsCreatorImpl(blockChainService)

	loadBalanceUcase := usecase.NewLoadBalanceUseCase(balance, opStorage)
	loadBalanceController := controllers.NewCliLoadBalanceController(loadBalanceUcase)

	showHelpPresenter := presenters.NewCliShowHelpPresenter(os.Stdout)
	showHelpUcase := usecase.NewShowHelpUseCase(showHelpPresenter)
	showHelpController := controllers.NewCliShowHelpController(showHelpUcase)

	depositPresenter := presenters.NewCliDepositPresenter(os.Stdout)
	depositUcase := usecase.NewDepositUseCase(balance, opStorage, opCreator, depositPresenter)
	depositController := controllers.NewCliDepositController(depositUcase)

	withdrawPresenter := presenters.NewCliWithdrawPresenter(os.Stdout)
	withdrawUcase := usecase.NewWithdrawUseCase(balance, opStorage, opCreator, withdrawPresenter)
	withdrawController := controllers.NewCliWithdrawController(withdrawUcase)

	showBalancePresenter := presenters.NewCliShowBalancePresenter(os.Stdout)
	showBalanceUcase := usecase.NewShowBalanceUseCase(balance, showBalancePresenter)
	showBalanceController := controllers.NewCliShowBalanceController(showBalanceUcase)

	showOpsPresenter := presenters.NewCliShowOperationsPresenter(os.Stdout)
	showOpsUcase := usecase.NewShowOperationsUseCase(opStorage, showOpsPresenter)
	showOpsController := controllers.NewCliShowOperationsController(showOpsUcase)

	undoLastPresenter := presenters.NewCliUndoLastPresenter(os.Stdout)
	undoLastUcase := usecase.NewUndoLastUseCase(balance, opStorage, undoLastPresenter)
	undoLastController := controllers.NewCliUndoLastController(undoLastUcase)

	commands := cli.Commands{
		cli.NewCommand(cli.LoadBalance, `^load$`, loadBalanceController),
		cli.NewCommand(cli.ShowHelpCommand, `^help$`, showHelpController),
		cli.NewCommand(cli.QuitCommand, `^quit$`, nil),
		cli.NewCommand(cli.DepositCommand, `^(deposit|dpt) (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, depositController),
		cli.NewCommand(cli.WithdrawCommand, `^(withdraw|wdw) (?P<currency>[a-zA-Z]{3}) (?P<amount>[0-9]+)$`, withdrawController),
		cli.NewCommand(cli.ShowBalanceCommand, `^(balance|bls)$`, showBalanceController),
		cli.NewCommand(cli.ShowOperationsCommand, `^(operations|ops)$`, showOpsController),
		cli.NewCommand(cli.UndoCommand, `^undo$`, undoLastController),
	}
	app := cli.NewCliApp(commands, cli.LoadBalance, cli.ShowHelpCommand)

	return app.Run()
}

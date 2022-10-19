package usecase

type ShowHelpUseCaseOutput interface {
	Present()
}

type ShowHelpUseCaseInput interface {
	Execute()
}

type ShowHelpUseCase struct {
	output ShowHelpUseCaseOutput
}

func NewShowHelpUseCase(output ShowHelpUseCaseOutput) *ShowHelpUseCase {
	return &ShowHelpUseCase{
		output: output,
	}
}

func (ucase *ShowHelpUseCase) Execute() {
	ucase.output.Present()
}

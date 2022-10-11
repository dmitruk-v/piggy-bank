package usecase

type InfoOperationsUseCase struct {
	opStorage OperationStorage
}

func NewInfoOperationsUseCase(opStorage OperationStorage) *InfoOperationsUseCase {
	return &InfoOperationsUseCase{
		opStorage: opStorage,
	}
}

func (ucase *InfoOperationsUseCase) Execute() error {
	return nil
}

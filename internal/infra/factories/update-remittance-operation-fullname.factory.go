package factories

import (
	"sync"
	"update-receivers/internal/application"
	"update-receivers/internal/infra/adapters"
)

type UpdateRemittanceOperationFullnameUseCaseFactory struct{}

func (nuroucf *UpdateRemittanceOperationFullnameUseCaseFactory) Generate(processList chan bool, wg *sync.WaitGroup) *application.UpdateRemittanceOperationFullnameUseCase {
	return application.NewUpdateRemittanceOperationFullnameUseCase(
		processList,
		wg,
		adapters.NewGetRemittanceOperationsFromJsonAdapter(),
		adapters.NewRequestHttpAdapter(),
	)
}

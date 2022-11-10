package application

import (
	"sync"
	"update-receivers/internal/application/ports"
	"update-receivers/internal/model"
)

type UpdateRemittanceOperationFullnameUseCase struct {
	processList                    chan bool
	wg                             *sync.WaitGroup
	getRemittanceOperationsAdapter ports.GetRemittanceOperationsOutboundPort
	requestHttpAdapter             ports.RequestHttpOutboundPort
}

func NewUpdateRemittanceOperationFullnameUseCase(
	processList chan bool,
	wg *sync.WaitGroup,
	getRemittanceOperationsFromJsonAdapter ports.GetRemittanceOperationsOutboundPort,
	requestHttpAdapter ports.RequestHttpOutboundPort,
) *UpdateRemittanceOperationFullnameUseCase {
	return &UpdateRemittanceOperationFullnameUseCase{
		processList:                    processList,
		wg:                             wg,
		getRemittanceOperationsAdapter: getRemittanceOperationsFromJsonAdapter,
		requestHttpAdapter:             requestHttpAdapter,
	}
}

func (uropfn *UpdateRemittanceOperationFullnameUseCase) Execute() {
	operations := uropfn.getRemittanceOperationsAdapter.GetOperations()
	for i, operation := range operations {
		uropfn.wg.Add(1)
		uropfn.processList <- true
		go func(processNumber int, wg *sync.WaitGroup, operation model.RemittanceOperation) {
			uropfn.requestHttpAdapter.Request(processNumber, operation)
			<-uropfn.processList
			wg.Done()
		}(i+1, uropfn.wg, operation)
	}
}

package main

import (
	"sync"
	"update-receivers/internal/infra/factories"
)

func main() {

	maxCallAtSameTime := 10
	processList := make(chan bool, maxCallAtSameTime)
	var wg sync.WaitGroup

	updateRemittanceFactory := &factories.UpdateRemittanceOperationFullnameUseCaseFactory{}
	updateRemittanceOperationFullnameUseCase := updateRemittanceFactory.Generate(
		processList,
		&wg,
	)

	updateRemittanceOperationFullnameUseCase.Execute()

	wg.Wait()
}

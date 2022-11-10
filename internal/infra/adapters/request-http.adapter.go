package adapters

import (
	"fmt"
	"strconv"
	"time"
	"update-receivers/internal/model"
)

type RequestHttpAdapter struct{}

func NewRequestHttpAdapter() *RequestHttpAdapter {
	return &RequestHttpAdapter{}
}

func (r *RequestHttpAdapter) Request(processNumber int, operation model.RemittanceOperation) {
	fmt.Println("Iniciando processo " + strconv.Itoa(processNumber) + "------------------")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------Finalizando processo " + strconv.Itoa(processNumber))
}

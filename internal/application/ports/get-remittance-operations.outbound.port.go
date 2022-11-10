package ports

import "update-receivers/internal/model"

type GetRemittanceOperationsOutboundPort interface {
	GetOperations() []model.RemittanceOperation
}

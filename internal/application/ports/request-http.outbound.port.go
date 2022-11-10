package ports

import "update-receivers/internal/model"

type RequestHttpOutboundPort interface {
	Request(processNumber int, operation model.RemittanceOperation)
}

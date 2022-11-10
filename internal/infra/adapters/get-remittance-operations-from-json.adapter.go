package adapters

import "update-receivers/internal/model"

type GetRemittanceOperationsFromJsonAdapter struct{}

func NewGetRemittanceOperationsFromJsonAdapter() *GetRemittanceOperationsFromJsonAdapter {
	return &GetRemittanceOperationsFromJsonAdapter{}
}

func (grop *GetRemittanceOperationsFromJsonAdapter) GetOperations() []model.RemittanceOperation {
	return []model.RemittanceOperation{
		{
			ID:             1,
			AccountNumber:  "123456",
			Iban:           "123456",
			AssignmentDate: "2020-01-01",
			CustomerID:     "123456",
			DocumentNumber: "123456",
			CompanyName:    "123456",
			IsIndividual:   true,
		},
		{
			ID:             2,
			AccountNumber:  "223456",
			Iban:           "223456",
			AssignmentDate: "2020-01-02",
			CustomerID:     "223456",
			DocumentNumber: "223456",
			CompanyName:    "223456",
			IsIndividual:   true,
		},
	}
}

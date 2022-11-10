package model

type RemittanceOperation = struct {
	ID             int    `json:"id"`
	AccountNumber  string `json:"account_number"`
	Iban           string `json:"iban"`
	AssignmentDate string `json:"assignment_date"`
	CustomerID     string `json:"customer_id"`
	DocumentNumber string `json:"document_number"`
	CompanyName    string `json:"company_name"`
	IsIndividual   bool   `json:"is_individual"`
}

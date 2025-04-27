package models

type FinancialTool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
	ServiceType string `json:"service_type"`
}

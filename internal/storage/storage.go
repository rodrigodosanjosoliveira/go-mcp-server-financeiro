package storage

import "go-mcp-server-financeiro/internal/models"

var financialTools []models.FinancialTool

func AddFinancialTool(tool models.FinancialTool) {
	financialTools = append(financialTools, tool)
}

func GetFinancialTools() []models.FinancialTool {
	return financialTools
}

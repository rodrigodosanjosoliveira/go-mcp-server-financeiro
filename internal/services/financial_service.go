package services

import (
	"go-mcp-server-financeiro/internal/models"
	"go-mcp-server-financeiro/internal/storage"
)

func RegisterFinancialTool(tool models.FinancialTool) {
	storage.AddFinancialTool(tool)
}

func ListFinancialTools() []models.FinancialTool {
	return storage.GetFinancialTools()
}

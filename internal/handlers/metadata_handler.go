package handlers

import (
	"encoding/json"
	"go-mcp-server-financeiro/internal/storage"
	"net/http"
)

func MetadataHandler(w http.ResponseWriter, r *http.Request) {
	type ToolMetadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ServiceType string `json:"service_type"`
	}

	var metadata []ToolMetadata

	tools := storage.GetFinancialTools()

	for _, tool := range tools {
		metadata = append(metadata, ToolMetadata{
			Name:        tool.Name,
			Description: tool.Description,
			ServiceType: tool.ServiceType,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"version": "1.0",
		"tools":   metadata,
	})
}

package handlers

import (
	"encoding/json"
	"go-mcp-server-financeiro/internal/models"
	"go-mcp-server-financeiro/internal/services"
	"net/http"
)

func RegisterFinancialToolHandler(w http.ResponseWriter, r *http.Request) {
	var tool models.FinancialTool
	err := json.NewDecoder(r.Body).Decode(&tool)
	if err != nil {
		http.Error(w, "Erro ao registrar ferramenta financeira", http.StatusBadRequest)
		return
	}
	services.RegisterFinancialTool(tool)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ferramenta financeira registrada com sucesso"})
}

func ListFinancialToolsHandler(w http.ResponseWriter, r *http.Request) {
	tools := services.ListFinancialTools()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tools)
}

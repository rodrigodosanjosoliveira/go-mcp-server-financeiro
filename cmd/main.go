package main

import (
	"log"
	"net/http"

	"go-mcp-server-financeiro/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado ou não carregado.")
	}

	r := mux.NewRouter()

	// MCP Server de Finanças
	r.HandleFunc("/register-financial-tool", handlers.RegisterFinancialToolHandler).Methods("POST")
	r.HandleFunc("/list-financial-tools", handlers.ListFinancialToolsHandler).Methods("GET")
	r.HandleFunc("/query-stock-price", handlers.QueryStockPriceHandler).Methods("POST")
	r.HandleFunc("/metadata", handlers.MetadataHandler).Methods("GET")

	log.Println("MCP Financial Server rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

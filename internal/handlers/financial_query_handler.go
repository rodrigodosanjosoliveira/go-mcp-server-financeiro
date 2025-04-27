package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var alphaVantageAPIKey = os.Getenv("ALPHA_VANTAGE_API_KEY")

func QueryStockPriceHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Symbol string `json:"symbol"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}

	apiUrl := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s",
		req.Symbol,
		alphaVantageAPIKey,
	)

	resp, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, "Erro ao consultar Alpha Vantage", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

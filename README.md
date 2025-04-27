# MCP Server Financeiro

Este projeto é um exemplo prático de como construir um MCP Server simples utilizando Go.  
Ele simula a gestão de ferramentas financeiras e realiza consultas reais de cotações de ações utilizando a API pública da Alpha Vantage.

O objetivo é mostrar como estruturar um servidor que siga os princípios básicos do Model Context Protocol (MCP), incluindo a exposição de um endpoint `/metadata`.

## Funcionalidades

- Registrar ferramentas financeiras dinamicamente
- Listar todas as ferramentas registradas
- Consultar preços de ações em tempo real (Alpha Vantage)
- Expor metadados no padrão MCP para agentes clientes

## Estrutura de Endpoints

### 1. Registrar ferramenta financeira

```
POST http://localhost:8080/register-financial-tool
Content-Type: application/json

{
  "name": "Consulta Ações",
  "description": "Consulta cotações de ações em tempo real",
  "endpoint": "https://www.alphavantage.co/query",
  "service_type": "cotacao"
}
```

### 2. Listar ferramentas financeiras

```
GET http://localhost:8080/list-financial-tools
Accept: application/json
```

### 3. Consultar preço de uma ação (Alpha Vantage)

```
POST http://localhost:8080/query-stock-price
Content-Type: application/json

{
  "symbol": "AAPL"
}
```

### 4. Consultar metadados do MCP Server

```
GET http://localhost:8080/metadata
Accept: application/json
```

## Pré-requisitos

- Conta gratuita na [Alpha Vantage](https://www.alphavantage.co/) para gerar sua API Key

## Como rodar o projeto

1. Clone o repositório

```
git clone https://github.com/seuusuario/go-mcp-server-financeiro.git
cd go-mcp-server-financeiro
```

2. Crie um arquivo `.env` na raiz do projeto:

```
ALPHA_VANTAGE_API_KEY=sua_api_key_aqui
```

3. Instale as dependências

```
go mod tidy
```

4. Rode o servidor

```
go run cmd/main.go
```

O servidor ficará disponível em `http://localhost:8080`.
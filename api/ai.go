package api

import (
	"encoding.json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	StatusCode int
	Balance int64
}

type Error struct {
	StatusCode int
	Message string
}
package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jim": {
		AuthToken: "123ABC",
		Username:  "jim",
	},
	"zen": {
		AuthToken: "123ABC",
		Username:  "zen",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1000,
		Username: "alex",
	},
	"jim": {
		Coins:    2000,
		Username: "jim",
	},
	"zen": {
		Coins:    3000,
		Username: "zen",
	},
}

func (db mockDB) GetUserLoginDetails(username string) *LoginDetails {

	// Simulate Database Call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db mockDB) GetUserCoins(username string) *CoinDetails {

	// Simmulate Database Call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db mockDB) SetupDatabase() error {
	return nil
}

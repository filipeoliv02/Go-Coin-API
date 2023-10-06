package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"filipe": LoginDetails{
		AuthToken: "1234567890",
		Username:  "filipeoliv02",
	},
	"john": LoginDetails{
		AuthToken: "0987654321",
		Username:  "johndoe",
	},
	"jane": LoginDetails{
		AuthToken: "qwertyuiop",
		Username:  "janedoe",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"filipe": CoinDetails{
		Coins:    9999,
		Username: "filipeoliv02",
	},
	"john": CoinDetails{
		Coins:    100,
		Username: "johndoe",
	},
	"jane": CoinDetails{
		Coins:    50,
		Username: "janedoe",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

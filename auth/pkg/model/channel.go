package model

type Channel struct {
	Base
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	API_KEY    string `json:"api_key" sql:",notnull"`
	APP_SECRET string `json:"app_secret"`
}

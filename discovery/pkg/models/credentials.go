package models

type IntegrationCredentials struct {
	Username string `json:"username"`
	APIKey   string `json:"api_key"`
	BaseURL  string `json:"base_url"`
}

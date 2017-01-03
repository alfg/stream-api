package models

type Client struct {
	ID        int64  `db:"id" json:"id,omitempty"`
	APIKey    string `db:"api_key" json:"api_key"`
	APISecret string `db:"api_secret" json:"api_secret"`
	Domain    string `db:"domain" json:"domain"`
	Email     string `db:"email" json:"email"`
}

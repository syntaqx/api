package model

type Secret struct {
	ID        string `json:"id"`
	Secret    string `json:"secret"`
	Used      bool   `json:"used"`
	CreatedAt int64  `json:"created_at"`
}

package model

type Weather struct {
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Description string `json:"description"`
}

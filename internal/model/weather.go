package model

import "net/http"

type Weather struct {
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Description string `json:"description"`
}

func (w *Weather) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

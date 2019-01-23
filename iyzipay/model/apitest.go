package model

type ApiTestResponse struct {
	Status     string `json:"status"`
	Locale     string `json:"locale"`
	SystemTime int64 `json:"systemTime"`
}
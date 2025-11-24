package models

type PageInfo struct {
	Key   string `json:"key"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

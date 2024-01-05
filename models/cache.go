package models

type CacheCreateResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

type CacheGetResponse struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

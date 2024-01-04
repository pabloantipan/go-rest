package models

type User struct {
	ID       int     `json:"id"`
	CacheID  string  `json:"cacheID"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

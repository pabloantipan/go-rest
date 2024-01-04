package repository

import (
	"context"
	"encoding/json"
	"strconv"

	"example.com/mod/models"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type CacheCreateResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

type CacheGetResponse struct {
	Status string      `json:"status"`
	User   models.User `json:"user"`
}

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{client: CacheConnect()}
}

func CacheConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis host and port
		Password: "",               // Set password if required
		DB:       0,                // Select Redis database
	})
}

func (rdb *CacheRepository) CreateUser(user models.User) CacheCreateResponse {
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	userID := uuid.New().String()

	err = rdb.client.Set(context.Background(), userID, userJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	return CacheCreateResponse{"User created successfully!", userID}
}

func (rdb *CacheRepository) GetUser(id int) CacheGetResponse {
	// Get the user by ID
	val, err := rdb.client.Get(context.Background(), strconv.Itoa(id)).Result()
	if err != nil {
		panic(err)
	}

	// Unmarshal user data from JSON
	var user models.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}

	return CacheGetResponse{"User found!", user}
}

package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"example.com/mod/models"
	"github.com/go-redis/redis/v8"
)

type UserCacheRepository struct {
	client *redis.Client
}

func NewCacheUserRepository() *UserCacheRepository {
	return &UserCacheRepository{client: CacheConnect()}
}

func CacheConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis host and port
		Password: "",               // Set password if required
		DB:       0,                // Select Redis database
	})
}

func (rdb *UserCacheRepository) CreateUser(user models.User) models.CacheCreateResponse {
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = rdb.client.Set(context.Background(), fmt.Sprintf("user:%s", user.CacheID), userJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	return models.CacheCreateResponse{Status: "User created successfully!"}
}

func (rdb *UserCacheRepository) GetByID(cacheID string) (models.User, error) {
	// Get the user by ID
	// val, err := rdb.client.Get(context.Background(), strconv.Itoa(id)).Result()
	val, err := rdb.client.Get(context.Background(), fmt.Sprintf("user:%s", cacheID)).Result()
	if err != nil {
		panic(err)
	}

	// Unmarshal user data from JSON
	var user models.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (rdb *UserCacheRepository) GetAll() ([]models.User, error) {
	var cursor uint64
	var keys []string
	for {
		var scanResult []string
		scanResult, cursor, err := rdb.client.Scan(context.Background(), cursor, "user:*", 10).Result()
		if err != nil {
			panic(err)
		}
		keys = append(keys, scanResult...)
		if cursor == 0 {
			break
		}
	}

	// Retrieve users for each key
	var users []models.User
	for _, key := range keys {
		userJSON, err := rdb.client.Get(context.Background(), key).Result()
		if err == nil {
			var user models.User
			err = json.Unmarshal([]byte(userJSON), &user)
			if err == nil {
				users = append(users, user)
			}
		}
	}

	return users, nil
}

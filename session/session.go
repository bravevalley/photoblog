package session

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RDB *redis.Client
var err error

func CheckLogin(loginID string) error {
	value, err := RDB.Exists(loginID).Result()
	if err != nil {
		log.Printf("Cant not read data from Redis: %v", err)
		return err
	}

	if value == 1 {
		return nil
	}

	return fmt.Errorf("ErrNoValue")
}

func CreateSession(key, value string) error {
	err = RDB.Set(key, value, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

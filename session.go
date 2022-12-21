package main

import (
	"fmt"
	"log"
	"time"
)

func checkLogin(loginID string) error {
	value, err := rdb.Exists(loginID).Result()
	if err != nil {
		log.Printf("Cant not read data from Redis: %v", err)
		return err
	}

	if value == 1 {
		return nil
	}

	return fmt.Errorf("ErrNoValue")
}

func createSession(key, value string) error {
	err = rdb.Set(key, value, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type FactoryOption[T any] func(*T)

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}

func RandomUUID() string {
	return uuid.New().String()
}

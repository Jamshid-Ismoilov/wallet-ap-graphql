package tools

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomGenerator() string {
	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(1000000)

	result := strconv.Itoa(randNum)

	return result
}

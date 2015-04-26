package util

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateRandomString() string {

	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int63()
	randStr := strconv.FormatInt(randInt, 36)

	return randStr
}

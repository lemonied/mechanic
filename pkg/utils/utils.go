package utils

import (
	"math/rand"
	"strconv"
	"time"
)

/*
RandomStr random string
*/
func RandomStr(length int) string {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := r.Int63()
	str := strconv.FormatInt(ran, 36)
  return str[0:length]
}

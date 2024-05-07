package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
  return min + r.Int63n(max-min+1)
}

func RandomString(length int) string {
  var sb strings.Builder
  k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string{
	return RandomString(6)
}

func RandomPassword() string{
	return RandomString(10)
}

func RandomEmail() string{
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

func RandomCurrency() string{
	currencies := []string{"USD","TWD","JPY","CNY","EUR"}
	n := len(currencies)
	return currencies[r.Intn(n)]
}
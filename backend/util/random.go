package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

 // 亂數函數
func RandomInt(min, max int64) int64 {
  return min + r.Int63n(max-min+1)
}

// 亂數字串函數
func RandomString(length int) string {
  var sb strings.Builder
  k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
//   亂數使用者名稱函數
func RandomUsername() string{
	return RandomString(6)
}

// 亂數密碼函數
func RandomPassword() string{
	return RandomString(10)
}

// 亂數信箱函數
func RandomEmail() string{
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

// 亂數貨幣函數
func RandomCurrency() string{
	currencies := []string{"USD","TWD","JPY","CNY","EUR"}
	n := len(currencies)
	return currencies[r.Intn(n)]
}
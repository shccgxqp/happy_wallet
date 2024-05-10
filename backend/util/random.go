package util

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/sqlc-dev/pqtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// 亂數Int 64函數
func RandomInt(min, max int64) int64 {
  return min + r.Int63n(max-min+1)
}

// 亂數生成float 64函數 小數點兩位
func RandomFloat(min, max float64) float64 {	
  return min + r.Float64()*(max-min)
}

// 亂數NullInt 64函數
func RandomNullInt(min, max int64) sql.NullInt64 {
	value := RandomInt(min, max)
	return sql.NullInt64{Int64: value, Valid: true}
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


// RandomTeamMembers 生成隨機團隊成員 JSON 資料
func RandomTeamMembers(num int) pqtype.NullRawMessage {
    members := make([]string, num)
    for i := 0; i < num; i++ {
        members[i] = RandomUsername()
    }
    data, err := json.Marshal(members)
    if err != nil {
        return pqtype.NullRawMessage{}
    }
    return pqtype.NullRawMessage{
        RawMessage: json.RawMessage(data),
        Valid:      true,
    }
}
package helper

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"math/rand"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

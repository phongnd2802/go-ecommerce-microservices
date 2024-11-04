package user

import (
	"fmt"
	"strings"
)


func GetUserKeyOTP(key string) string {
	return fmt.Sprintf("user:%s:otp", key)
}

func GetNicknameFromEmail(email string) string {
	split := strings.Split(email, "@")
	return strings.ToLower(split[0])
}

package user

import "fmt"


func GetUserKeyOTP(key string) string {
	return fmt.Sprintf("user:%s:otp", key)
}
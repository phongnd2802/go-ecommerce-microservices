package crypto

import "golang.org/x/crypto/bcrypt"

func HashPasswordWithSalt(password string, salt string) (string, error) {
	passwordSalt := password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func ComparePasswordWithSalt(password string, salt string, hashedPassword string) bool {
	passwordSalt := password + salt
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordSalt)) == nil
}


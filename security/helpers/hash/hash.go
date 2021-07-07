package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPw string) (string, error) {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(plainPw), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", hashErr
	}

	return string(hash), hashErr
}

func CompareHashedPwWithPlainPw(hashedPassword string, plainPassword string) bool {
	byteHashPw := []byte(hashedPassword)
	bytePlainPw := []byte(plainPassword)

	compErr := bcrypt.CompareHashAndPassword(byteHashPw, bytePlainPw)
	if compErr != nil {
		return false
	}
	return true
}

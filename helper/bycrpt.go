package helper

import "golang.org/x/crypto/bcrypt"

func Hashpw(pw string) string {
	PwHashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(PwHashed)
}

func ValidatePw(password, hashedPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password))

	if err != nil {
		return false
	} else {
		return true
	}

}

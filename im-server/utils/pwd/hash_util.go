package pwd

import "golang.org/x/crypto/bcrypt"

func HashPwd(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(password)
}

func CheckPwd(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

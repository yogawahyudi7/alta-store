package controllers

import "golang.org/x/crypto/bcrypt"

//================
//HASH PASSWORD
//================
func Hashpwd(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes), err
}

func Checkpwd(hash, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}

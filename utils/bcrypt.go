package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd []byte)(string,error){
	hash,err:=bcrypt.GenerateFromPassword(pwd,bcrypt.MinCost)
	if err!=nil{
		return "",err
	}
	return string(hash),nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
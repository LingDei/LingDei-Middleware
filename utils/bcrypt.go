package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 密码Bcrypt加密
func GetPwdBcryptHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// ValidatePwdBcryptHash 密码校验
func ValidatePwdBcryptHash(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
}

// StringToSlice 将字符串转换为切片，用逗号分隔
func StringToSlice(str string) []string {
	slice := strings.Split(str, ",")
	return slice
}

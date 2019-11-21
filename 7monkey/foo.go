package _monkey

import (
	"strconv"
)

func MysqlFunc(age int) string {
	// some filtering on age
	user := "c" + strconv.Itoa(age)
	return user
}

func WrapMysqlFunc(age int) string {
	user := MysqlFunc(age)
	user = strconv.Itoa(age) + user
	return user
}

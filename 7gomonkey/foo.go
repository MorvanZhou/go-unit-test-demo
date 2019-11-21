package _gomonkey

import (
	"strconv"
)

func MysqlFunc(age int) (users []string) {
	// some filtering on age
	users = []string{"c" + strconv.Itoa(age), "f" + strconv.Itoa(age)}
	return users
}

func WrapMysqlFunc(age int) []string {
	users := MysqlFunc(age)
	for i, u := range users {
		users[i] = strconv.Itoa(age) + u
	}
	return users
}

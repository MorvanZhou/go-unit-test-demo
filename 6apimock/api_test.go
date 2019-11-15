package _apimock_test

import (
	_apimock "github.com/morvanzhou/unittest-demo/6apimock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetApiAddOne(t *testing.T) {
	defer gock.Off()

	gock.New("http://your-api.com").
		Get("/get"). // get 方法
		Reply(200).
		JSON(map[string]int{"Return": 1}) // 回复 1

	res := _apimock.GetApiAddOne()

	assert.Equal(t, res, 2) // 1 + 1 = 2

	assert.True(t, gock.IsDone())
}

func TestPostApiAddOne(t *testing.T) {
	defer gock.Off()

	gock.New("http://your-api.com").
		Post("/post"). // post 方法
		MatchType("json").
		JSON(map[string]int{"Param": 1}). // 请求为 1 时
		Reply(200).
		JSON(map[string]int{"Return": 1}) // 回复 1

	res := _apimock.PostApiAddOne(1) // 发送请求 1
	assert.Equal(t, res, 2)          // 1 + 1 = 2

	gock.New("http://your-api.com").
		Post("/post"). // post 方法
		MatchType("json").
		JSON(map[string]int{"Param": 2}). // 请求为 2 时
		Reply(200).
		JSON(map[string]int{"Return": 2}) // 回复 2

	res = _apimock.PostApiAddOne(2) // 发送请求 2
	assert.Equal(t, res, 3)         // 2 + 1 = 3

	assert.True(t, gock.IsDone())
}

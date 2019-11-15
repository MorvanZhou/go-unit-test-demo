package _apimock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type getResp struct {
	Return int
}

func GetApiAddOne() int {
	resp, err := http.Get("http://your-api.com/get")
	if err != nil {
		return 0
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var resp_ getResp
	if err := json.Unmarshal(body, &resp_); err != nil {
		return 0
	}
	return resp_.Return + 1
}

type postReq struct {
	Param int
}
type postResp struct {
	Return int
}

func PostApiAddOne(x int) int {
	query := &postReq{Param: x}
	b, _ := json.Marshal(query)

	resp, err := http.Post("http://your-api.com/post", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return 0
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var resp_ postResp
	if err := json.Unmarshal(body, &resp_); err != nil {
		return 0
	}
	return resp_.Return + 1
}

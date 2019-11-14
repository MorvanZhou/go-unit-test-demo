package _gin_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/morvanzhou/unittest-demo/4gin"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHelloWorld(t *testing.T) {
	tt := []struct {
		name string
		in string
		expected string
	}{
		{"test1", "human", "human"},
		{"test2", "dog", "dog"},
	}

	// Grab our router
	router := _gin.SetupRouter()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			// request
			reqStr := fmt.Sprintf(`{"Param": "%s"}`, tc.in)
			r := strings.NewReader(reqStr)

			// get response
			w := performRequest(router, "POST", "/post_test", r)
			assert.Equal(t, http.StatusOK, w.Code)

			// verify return
			var response map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Nil(t, err)

			value, exists := response["Return"]
			assert.True(t, exists)
			assert.Equal(t, tc.expected, value)
		})
	}
}

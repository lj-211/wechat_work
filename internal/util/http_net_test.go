package util

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	data := struct{ Val int }{Val: 1234}
	data_str, _ := json.Marshal(data)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		// 在这里mock需要的数据
		w.WriteHeader(200)
		_ = io.WriteString(w, string(data_str))
	}))
	defer server.Close()

	body_str, req_err := Post(server.URL, "", "application/json;charset=UTF-8")
	assert.Equal(t, string(data_str), body_str)
	assert.Nil(t, req_err)
}

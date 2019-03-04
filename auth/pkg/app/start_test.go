package app

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/alamin-mahamud/arya/auth/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	r := httptest.NewRequest("GET", "http://localhost"+":8080"+"/ping", nil)
	w := httptest.NewRecorder()
	handler.Ping(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEmpty(t, resp.Header.Get("Content-Type"))
	assert.NotEmpty(t, string(body))
}

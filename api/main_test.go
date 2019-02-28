package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/alamin-mahamud/arya/api/pkg"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	r := httptest.NewRequest("GET", "http://localhost"+":8080"+"/ping", nil)
	w := httptest.NewRecorder()
	pkg.Ping(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEmpty(t, resp.Header.Get("Content-Type"))
	assert.NotEmpty(t, string(body))
}

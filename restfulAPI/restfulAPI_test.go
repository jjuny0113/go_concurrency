package restfulAPI

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/students", nil)

	mux := makeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)

}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := makeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/students/1", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := makeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/students/3", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)

}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := makeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/students/1", nil)

	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)
}

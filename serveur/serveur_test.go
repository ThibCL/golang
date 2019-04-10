package serveur

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ThibCL/gotest/store"

	"github.com/stretchr/testify/assert"
)

func createBody(lang string, hello string) *bytes.Buffer {
	var body addLanguageRequest
	body.Lang = lang
	body.Hello = hello
	buf, _ := json.Marshal(body)
	buff := bytes.NewBuffer(buf)
	return buff
}

func TestAddHello(t *testing.T) {
	InitializeStore()
	buf := createBody("en", "hello")

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()
	AddHello(res, req)
	resp := res.Result()

	assert.Equal(t, 200, resp.StatusCode)

}

func TestAddHelloStoreFunctionErr(t *testing.T) {
	InitializeStore()
	s.AddLang("en", "Hello")
	buf := createBody("En", "hello")

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()
	AddHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, store.ErrAlreadyExists.Error()+"\n", string(body))
}

func TestAddHelloValidationErr(t *testing.T) {
	InitializeStore()
	buf := createBody("Eng", "hello")

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()
	AddHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, ErrWrongFormat.Error()+"\n", string(body))
}

func TestAddHelloErrBody(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", strings.NewReader("test"))
	res := httptest.NewRecorder()
	AddHello(res, req)
	resp := res.Result()

	assert.Equal(t, 400, resp.StatusCode)

}

func TestSayHello(t *testing.T) {
	InitializeStore()
	s.AddLang("en", "hello")

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()
	SayHello(res, req)
	resp := res.Result()

	assert.Equal(t, 200, resp.StatusCode)

}
func TestSayHelloParamMissing(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("GET", "http://localhost:9000/hello", nil)
	res := httptest.NewRecorder()
	SayHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Param 'lang' missing"+"\n", string(body))
}

func TestSayHelloStoreFunctionErr(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()
	SayHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, store.ErrNotKnown.Error()+"\n", string(body))
}

func TestSayHelloValidationErr(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=eng", nil)
	res := httptest.NewRecorder()
	SayHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, ErrWrongFormat.Error()+"\n", string(body))
}

func TestDeleteHello(t *testing.T) {
	InitializeStore()
	s.AddLang("en", "hello")

	req := httptest.NewRequest("DELETE", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()
	DeleteHello(res, req)
	resp := res.Result()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteHelloValidationErr(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("DELETE", "http://localhost:9000/hello?lang=eng", nil)
	res := httptest.NewRecorder()
	DeleteHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, ErrWrongFormat.Error()+"\n", string(body))
}

func TestDeleteHelloParamMissing(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("DELETE", "http://localhost:9000/hello", nil)
	res := httptest.NewRecorder()
	DeleteHello(res, req)
	resp := res.Result()

	assert.Equal(t, 400, resp.StatusCode)

}

func TestDeleteHelloStoreFunctionErr(t *testing.T) {
	InitializeStore()

	req := httptest.NewRequest("DELETE", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()
	DeleteHello(res, req)
	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, store.ErrNotKnown.Error()+"\n", string(body))
}

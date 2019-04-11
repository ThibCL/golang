package serveur

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ThibCL/gotest/serveur/mocks"
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

	buf := createBody("en", "hello")
	str := new(mocks.HelloStore)

	str.On("AddLang", "en", "hello").Return(nil)

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()
	helloService := HelloService{str}
	helloService.AddHello(res, req)
	resp := res.Result()

	assert.Equal(t, 200, resp.StatusCode)
	str.AssertExpectations(t)

}

func TestAddHelloStoreFunctionErr(t *testing.T) {

	buf := createBody("En", "hello")
	str := new(mocks.HelloStore)
	str.On("AddLang", "en", "hello").Return(errors.New("Already exists"))

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.AddHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertExpectations(t)
}

func TestAddHellovalidateLangErr(t *testing.T) {

	buf := createBody("Eng", "hello")
	str := new(mocks.HelloStore)

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", buf)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.AddHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "AddLang")
}

func TestAddHelloErrBody(t *testing.T) {

	str := new(mocks.HelloStore)

	req := httptest.NewRequest("POST", "http://localhost:9000/hello", strings.NewReader("test"))
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.AddHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "AddLang")

}

func TestSayHello(t *testing.T) {
	str := new(mocks.HelloStore)
	str.On("Hello", "en").Return("Hello", nil)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.SayHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	str.AssertExpectations(t)

}

func TestSayHelloParamMissing(t *testing.T) {
	str := new(mocks.HelloStore)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.SayHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "Hello")
}

func TestSayHelloStoreFunctionErr(t *testing.T) {
	str := new(mocks.HelloStore)
	str.On("Hello", "en").Return("", errors.New("Language not Known"))

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.SayHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "Hello")
}

func TestSayHellovalidateLangErr(t *testing.T) {
	str := new(mocks.HelloStore)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=eng", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.SayHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "Hello")
}

func TestDeleteHello(t *testing.T) {
	str := new(mocks.HelloStore)
	str.On("DeleteLang", "en").Return(nil)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.DeleteHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	str.AssertExpectations(t)
}

func TestDeleteHellovalidateLangErr(t *testing.T) {
	str := new(mocks.HelloStore)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=eng", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.DeleteHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "DeleteLang")
}

func TestDeleteHelloParamMissing(t *testing.T) {
	str := new(mocks.HelloStore)

	req := httptest.NewRequest("GET", "http://localhost:9000/hello", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.DeleteHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertNotCalled(t, "DeleteLang")

}

func TestDeleteHelloStoreFunctionErr(t *testing.T) {
	str := new(mocks.HelloStore)
	str.On("DeleteLang", "en").Return(errors.New("Language Not Known"))

	req := httptest.NewRequest("GET", "http://localhost:9000/hello?lang=en", nil)
	res := httptest.NewRecorder()

	helloService := HelloService{str}
	helloService.DeleteHello(res, req)
	resp := res.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 400, resp.StatusCode)
	str.AssertExpectations(t)
}

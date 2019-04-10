//Package serveur ...
package serveur

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ThibCL/gotest/store"
)

type addLanguageRequest struct {
	Lang  string `json:"lang"`
	Hello string `json:"hello"`
}

var s store.Store

//SayHello d
func SayHello(res http.ResponseWriter, req *http.Request) {

	lang, langExist := req.URL.Query()["lang"]
	if !langExist {
		http.Error(res, "Param 'lang' missing", http.StatusBadRequest)
		return
	}
	err := validation(lang[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	hello, err := s.Hello(lang[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(res, hello)

}

//AddHello df
func AddHello(res http.ResponseWriter, req *http.Request) {
	var newLang addLanguageRequest
	body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(body, &newLang)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = validation(newLang.Lang)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddLang(newLang.Lang, newLang.Hello)
	if err != nil {

		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	bodyResp, _ := json.Marshal("Language Added")
	res.Write(bodyResp)

}

//DeleteHello dsf
func DeleteHello(res http.ResponseWriter, req *http.Request) {

	lang, langExist := req.URL.Query()["lang"]
	if !langExist {
		http.Error(res, "Param 'lang' missing", http.StatusBadRequest)
		return
	}

	err := validation(lang[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.DeleteLang(lang[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(res, "Language deleted")
}

//InitializeStore initialize the store
func InitializeStore() {
	s = store.NewStore()
}

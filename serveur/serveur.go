package serveur

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//HelloStore f
type HelloStore interface {
	Hello(string) (string, error)
	DeleteLang(string) error
	AddLang(string, string) error
}

//HelloService f
type HelloService struct {
	str HelloStore
}

type addLanguageRequest struct {
	Lang  string `json:"lang"`
	Hello string `json:"hello"`
}

type bodyResp struct {
	Text string `json:"text"`
}

//NewHelloService dsf
func NewHelloService(s HelloStore) *HelloService {
	return &HelloService{str: s}
}

//Register dfs
func (s *HelloService) Register(r *mux.Router) {
	r.HandleFunc("/hello", s.AddHello).Methods("POST")
}

//AddHello : Add a new language in the store
func (s *HelloService) AddHello(res http.ResponseWriter, req *http.Request) {
	var newLang addLanguageRequest
	body, _ := ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(body, &newLang); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateLang(&newLang.Lang); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.str.AddLang(newLang.Lang, newLang.Hello); err != nil {

		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	bodyResp, _ := json.Marshal(bodyResp{"Language added"})
	res.Write(bodyResp)

}

//Deleter fds
func (s *HelloService) Deleter(r *mux.Router) {
	r.HandleFunc("/hello", s.DeleteHello).Methods("DELETE")
}

//DeleteHello : Delete a language of a store
func (s *HelloService) DeleteHello(res http.ResponseWriter, req *http.Request) {

	lang, langExist := req.URL.Query()["lang"]
	if !langExist {
		http.Error(res, "Param 'lang' missing", http.StatusBadRequest)
		return
	}

	if err := ValidateLang(&lang[0]); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.str.DeleteLang(lang[0]); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	bodyResp, _ := json.Marshal(bodyResp{"Language deleted"})
	res.Write(bodyResp)
}

//Sayer sfdf
func (s *HelloService) Sayer(r *mux.Router) {
	r.HandleFunc("/hello", s.SayHello).Methods("GET")
}

//SayHello : return the traduction of hello in the language asked
func (s *HelloService) SayHello(res http.ResponseWriter, req *http.Request) {

	lang, langExist := req.URL.Query()["lang"]
	if !langExist {
		http.Error(res, "Param 'lang' missing", http.StatusBadRequest)
		return
	}

	if err := ValidateLang(&lang[0]); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	hello, err := s.str.Hello(lang[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	bodyResp, _ := json.Marshal(bodyResp{hello})
	res.Write(bodyResp)

}

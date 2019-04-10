//Package store ...
package store

import (
	"errors"
	"strings"
)

//Store : Object that store all the language supported by the api
type Store struct {
	lang map[string]string
}

//ErrNotKnown dsfq
var ErrNotKnown = errors.New("Language not known")

//ErrAlreadyExists fds
var ErrAlreadyExists = errors.New("Language already exists")

//NewStore : Constructor for the Store Object
func NewStore() Store {
	s := Store{lang: make(map[string]string)}
	return s
}

//Hello : method to get hello in an language
func (str *Store) Hello(lang string) (string, error) {
	lang = strings.ToLower(lang)
	_, exist := str.lang[lang]
	if !exist {
		return "", ErrNotKnown
	}

	return str.lang[lang], nil

}

//AddLang : Method to add language to the store
func (str *Store) AddLang(lang string, hello string) error {

	lang = strings.ToLower(lang)
	_, exist := str.lang[lang]
	if exist {
		return ErrAlreadyExists
	}

	str.lang[lang] = hello
	return nil
}

//DeleteLang : Method to delete one of the language of the store
func (str *Store) DeleteLang(lang string) error {

	lang = strings.ToLower(lang)
	_, exist := str.lang[lang]
	if !exist {
		return ErrNotKnown
	}

	delete(str.lang, lang)
	return nil
}

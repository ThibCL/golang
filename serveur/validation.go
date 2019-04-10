package serveur

import (
	"errors"

	"golang.org/x/text/language"
)

//ErrDoesNotExist ds
var ErrDoesNotExist = errors.New("Language does not exist")

//ErrWrongFormat sfd
var ErrWrongFormat = errors.New("Language should be two letter")

func validation(lang string) error {

	if len(lang) != 2 {
		return ErrWrongFormat
	}

	_, err := language.ParseBase(lang)
	if err != nil {
		return ErrDoesNotExist
	}
	return nil

}

package serveur

import (
	"errors"
	"strings"

	"golang.org/x/text/language"
)

//ValidateLang d
func ValidateLang(lang *string) error {

	if *lang = strings.ToLower(*lang); len(*lang) != 2 {
		return errors.New("Language should be two letter")
	}

	if _, err := language.ParseBase(*lang); err != nil {
		return errors.New("Language does not exist")
	}
	return nil

}

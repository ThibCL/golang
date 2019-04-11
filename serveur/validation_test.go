package serveur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLang(t *testing.T) {
	lang := "en"
	err := ValidateLang(&lang)
	assert.Nil(t, err)
}

func TestValidateLangErrWrongFormat(t *testing.T) {
	lang := "eng"
	err := ValidateLang(&lang)
	assert.EqualError(t, err, "Language should be two letter")
}

func TestValidateLangErrDoesNotExist(t *testing.T) {
	lang := "xx"
	err := ValidateLang(&lang)
	assert.EqualError(t, err, "Language does not exist")
}

package serveur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	err := validation("en")
	assert.Nil(t, err)
}

func TestValidationErrWrongFormat(t *testing.T) {
	err := validation("eng")
	assert.EqualError(t, err, ErrWrongFormat.Error())
}

func TestValidationErrDoesNotExist(t *testing.T) {
	err := validation("xx")
	assert.EqualError(t, err, ErrDoesNotExist.Error())
}

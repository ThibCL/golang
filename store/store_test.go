package store

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testpair struct {
	lang  string
	hello string
}

// func TestAddLangErrWrongFormat(t *testing.T) {
// 	s := NewStore()
// 	err := s.AddLang("Eng", "hello")
// 	assert.EqualError(t, err, ErrWrongFormat.Error())
// }

// func TestAddLangErrDoesNotExists(t *testing.T) {
// 	s := NewStore()

// 	err := s.AddLang("xx", "hi")
// 	assert.EqualError(t, err, ErrDoesNotExists.Error())
// }

func TestAddLang(t *testing.T) {
	s := NewStore()

	tests := []testpair{
		{"en", "Hello"},
		{"it", "Buongiorno"},
		{"fr", "Bonjour"},
	}

	for _, pair := range tests {
		assert.Equal(t, nil, s.AddLang(pair.lang, pair.hello))
		assert.Equal(t, pair.hello, s.lang[strings.ToLower(pair.lang)])
	}
}

func TestAddLangErrAlreadyExists(t *testing.T) {
	s := NewStore()
	s.AddLang("En", "Hello")
	err := s.AddLang("En", "Hello")
	assert.EqualError(t, err, ErrAlreadyExists.Error())

}

// func TestDeleteLangErrWrongFormat(t *testing.T) {
// 	s := NewStore()

// 	err := s.DeleteLang("Eng")
// 	assert.EqualError(t, err, ErrWrongFormat.Error())
// }

func TestDeleteLang(t *testing.T) {
	s := NewStore()
	s.AddLang("En", "Hello")

	err := s.DeleteLang("En")
	assert.Empty(t, s.lang["En"])
	assert.Equal(t, nil, err)
}

func TestDeleteLangErrNotKnown(t *testing.T) {
	s := NewStore()

	err := s.DeleteLang("En")
	assert.EqualError(t, err, ErrNotKnown.Error())
}

// func TestDeleteLangErrDoesNotExists(t *testing.T) {
// 	s := NewStore()

// 	err := s.DeleteLang("xx")
// 	assert.EqualError(t, err, ErrDoesNotExists.Error())
// }

// func TestHelloErrDoesNotExists(t *testing.T) {
// 	s := NewStore()

// 	hello, err := s.Hello("xx")
// 	assert.EqualError(t, err, ErrDoesNotExists.Error())
// 	assert.Equal(t, "", hello)
// }

func TestHelloErrNotKnown(t *testing.T) {
	s := NewStore()

	hello, err := s.Hello("En")
	assert.EqualError(t, err, ErrNotKnown.Error())
	assert.Equal(t, "", hello)
}

// func TestHelloErrWrongFormat(t *testing.T) {
// 	s := NewStore()

// 	hello, err := s.Hello("Eng")
// 	assert.EqualError(t, err, ErrWrongFormat.Error())
// 	assert.Equal(t, "", hello)
// }

func TestHello(t *testing.T) {
	s := NewStore()

	s.AddLang("en", "Hello")
	hello, err := s.Hello("en")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Hello", hello)
}

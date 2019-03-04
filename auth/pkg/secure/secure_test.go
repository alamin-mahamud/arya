package secure_test

import (
	"crypto/sha1"
	"testing"

	"github.com/alamin-mahamud/arya/auth/pkg/secure"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	testCases := []struct {
		name          string
		requiredInput string
		inputs        []string
		want          bool
	}{
		{
			name:          "Insecure Password",
			requiredInput: "notSec",
			want:          false,
		},
		{
			name:          "Password Matches Input Fields",
			requiredInput: "johndoe92",
			inputs:        []string{"john", "Doe"},
			want:          false,
		},
		{
			name:          "Secure Passwords",
			requiredInput: "cakkgophers",
			inputs:        []string{"john", "Doe"},
			want:          true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := secure.New(1, nil)
			got := s.Password(tt.requiredInput, tt.inputs...)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestHashAndMatch(t *testing.T) {
	testCases := []struct {
		name string
		pass string
		want bool
	}{
		{
			name: "Success",
			pass: "gamepad",
			want: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := secure.New(1, nil)
			hash := s.Hash(tt.pass)
			assert.Equal(t, tt.want, s.HashMatchesPassword(hash, tt.pass))
		})
	}
}

func TestToken(t *testing.T) {
	s := secure.New(1, sha1.New())
	token := "token"
	newToken := s.Token(token)
	assert.NotEqual(t, newToken, token)
}

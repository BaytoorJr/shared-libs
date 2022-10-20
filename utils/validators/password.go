package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"unicode"
)

// PasswordValidate
// password validator
func PasswordValidate(password string) error {
	var (
		hasUppercase bool
		hasLowercase bool
		hasNumber    bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUppercase {
		return errors.InvalidCharacter.SetDevMessage("invalid password: missing uppercase chars")
	}

	if !hasLowercase {
		return errors.InvalidCharacter.SetDevMessage("invalid password: missing lowercase chars")
	}

	if !hasNumber {
		return errors.InvalidCharacter.SetDevMessage("invalid password: missing numbers")
	}

	if len(password) < 8 || len(password) > 64 {
		return errors.InvalidCharacter.SetDevMessage("invalid password length: 8 - 64 needed")
	}

	return nil
}

package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"unicode"
)

const langLength = 2

// LangValidate
// lang validator
func LangValidate(lang string) error {
	if len(lang) != langLength {
		return errors.InvalidCharacter.SetDevMessage("lang length must be 2 letters")
	}

	if !unicode.IsLower(rune(lang[0])) || !unicode.IsLower(rune(lang[1])) {
		return errors.InvalidCharacter.SetDevMessage("lang must be lowercase")
	}

	if !unicode.IsLetter(rune(lang[0])) || !unicode.IsLetter(rune(lang[1])) {
		return errors.InvalidCharacter.SetDevMessage("lang must have only letters")
	}

	return nil
}

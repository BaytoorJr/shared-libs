package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"strings"
	"unicode"
)

// StringNonEmptyValidate
// string non-empty validator
func StringNonEmptyValidate(field, str string) error {
	if str == "" {
		return errors.InvalidCharacter.SetDevMessage(field + " is empty")
	}

	return nil
}

// StringLengthValidate
// string length validator
func StringLengthValidate(field, str string, length int) error {
	if len(str) != length {
		return errors.InvalidCharacter.SetDevMessage("invalid length of " + field)
	}

	return nil
}

// StringSlugValidate
// string slug validator
func StringSlugValidate(field, str string) error {
	if str == "" {
		return errors.InvalidCharacter.SetDevMessage("invalid length of " + field)
	}

	for _, symbol := range str {
		if (!unicode.IsLetter(symbol) || !unicode.IsLower(symbol)) &&
			!strings.ContainsAny(string(symbol), "-_") {
			return errors.InvalidCharacter.SetDevMessage("invalid characters in " + field)
		}
	}

	return nil
}

// StringMatchOne
// string match validator
func StringMatchOne(field, str string, arr []string) error {
	for i := 0; i < len(arr); i++ {
		if str == arr[i] {
			return nil
		}
	}

	return errors.InvalidCharacter.SetDevMessage("empty or invalid value in " + field)
}

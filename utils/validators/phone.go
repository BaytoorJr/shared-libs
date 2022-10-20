package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"strconv"
)

const phoneLength = 11

// PhoneNumberValidate
// phone numbers validator
func PhoneNumberValidate(phone string) error {
	if len(phone) != phoneLength {
		return errors.InvalidCharacter.SetDevMessage("invalid phone length")
	}

	_, err := strconv.Atoi(phone)
	if err != nil {
		return errors.InvalidCharacter.SetDevMessage("invalid phone provided")
	}

	return nil
}

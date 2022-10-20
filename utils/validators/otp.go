package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"unicode"
)

const otpLength = 4

// OTPValidate
// otp validator
func OTPValidate(otp string) error {
	err := errors.InvalidCharacter.SetDevMessage("invalid otp")

	if len(otp) != otpLength {
		return err
	}

	for _, c := range otp {
		if !unicode.IsDigit(c) {
			return err
		}
	}

	return nil
}

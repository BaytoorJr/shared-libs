package validators

import (
	"github.com/BaytoorJr/shared-libs/errors"
	"github.com/google/uuid"
)

// UUIDValidate
// uuid validator
func UUIDValidate(field, uid string) error {
	_, err := uuid.Parse(uid)
	if err != nil {
		return errors.InvalidCharacter.SetDevMessage("invalid " + field)
	}

	return nil
}

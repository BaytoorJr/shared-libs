package auth

import (
	"encoding/json"
	"fmt"
	"github.com/BaytoorJr/shared-libs/errors"
	"github.com/BaytoorJr/shared-libs/utils/validators"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

// Validate token & get user data
func getUserData(token *jwt.Token) (*UserData, error) {
	userData := new(UserData)
	meta := make(map[string]MetaData)

	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	// Set User ID
	userID := fmt.Sprintf("%v", claims["user_id"])
	err = validators.UUIDValidate("user_id", userID)
	if err != nil {
		return nil, err
	}

	userData.UserID = userID

	// Set User Mobile
	userData.UserMobile = fmt.Sprintf("%v", claims["mobile"])

	// Set Super Admin flag
	isSuperAdmin, err := strconv.ParseBool(fmt.Sprintf("%v", claims['is_superadmin']))
	if err != nil {
		return nil, errors.SerializeError.SetDevMessage(err.Error())
	}
	userData.IsSuperAdmin = isSuperAdmin

	// Get base profile ID
	if claims["base_profile_id"] != nil {
		baseProfileID := fmt.Sprintf("%v", claims["base_profile_id"])
		err = validators.UUIDValidate("base_profile_id", baseProfileID)
		if err != nil {
			return nil, err
		}
		userData.ProfileID = &baseProfileID
	}

	// Read meta data
	err = json.Unmarshal([]byte(claims["meta"].(string)), &meta)
	if err != nil {
		return nil, errors.SerializeError.SetDevMessage(err.Error())
	}

	projectMeta, ok := meta[config.ProjectSlug]
	if !ok {
		return userData, nil
	}

	userData.ProfileID = &projectMeta.ProfileID
	userData.RoleID = projectMeta.RoleID

	return userData, nil
}

// Check user permissions
func checkUserPermissions(roleID *int, actions, optionalActions []string) (bool, error) {
	var hasOptional bool

	if roleID == nil {
		if len(actions) > 0 {
			return false, errors.InvalidCharacter.SetDevMessage("role not presented in JWT")
		}

		return false, nil
	}

	actionsMap, err := getRoleActions(*roleID)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(actions); i++ {
		_, ok := actionsMap[actions[i]]
		if !ok {
			return false, errors.AccessDenied.SetDevMessage("don't have permission to access")
		}
	}

	if len(optionalActions) > 0 {
		hasOptional = true

		for i := 0; i < len(optionalActions); i++ {
			_, ok := actionsMap[optionalActions[i]]
			if !ok {
				hasOptional = false
			}
		}
	}

	return hasOptional, nil
}

// Check user permissions
func checkUserPermissions(roleID *int, actions, optionalActions []string) (bool, error) {
	var hasOptional bool

	if roleID == nil {
		if len(actions) > 0 {
			return false, errors.InvalidCharacter.SetDevMessage("role not presented in JWT")
		}

		return false, nil
	}

	actionsMap, err := getRoleActions(*roleID)
	if err != nil {
		return false, err
	}

	for i := 0; i < len(actions); i++ {
		_, ok := actionsMap[actions[i]]
		if !ok {
			return false, errors.AccessDenied.SetDevMessage("don't have permissions to access")
		}
	}

	if len(optionalActions) > 0 {
		hasOptional = true

		for i := 0; i < len(optionalActions); i++ {
			_, ok := actionsMap[optionalActions[i]]
			if !ok {
				hasOptional = false
			}
		}
	}

	return hasOptional, nil
}

// Check user optional permissions
func checkUserOptionalPermissions(roleID *int, optionalActions []string) ([]string, error) {
	var availableActions []string

	if roleID == nil {
		return nil, nil
	}

	actionsMap, err := getRoleActions(*roleID)
	if err != nil {
		return nil, err
	}

	if len(optionalActions) > 0 {
		for i := 0; i < len(optionalActions); i++ {
			_, ok := actionsMap[optionalActions[i]]
			if ok {
				availableActions = append(availableActions, optionalActions[i])
			}
		}
	}

	return availableActions, nil
}

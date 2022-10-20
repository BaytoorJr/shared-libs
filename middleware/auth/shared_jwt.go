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
func getSharedUserData(token *jwt.Token) (*UserData, error) {
	userData := new(UserData)
	meta := make(map[string]MetaData)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.InvalidCharacter.SetDevMessage("token invalid")
	}

	// Set user ID
	userID := fmt.Sprintf("%v", claims["user_id"])
	err := validators.UUIDValidate("user_id", userID)
	if err != nil {
		return nil, err
	}
	userData.UserID = userID

	// Set user mobile
	userData.UserMobile = fmt.Sprintf("%v", claims["mobile"])

	// Set super admin flag
	isSuperAdmin, err := strconv.ParseBool(fmt.Sprintf("%v", claims["is_superadmin"]))
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

	userData.Meta = meta

	return userData, nil
}

// Check user permissions
func checkSharedUserPermissions(meta map[string]MetaData, actions []string) (*[]string, error) {
	var projects []string

	if len(meta) == 0 {
		if len(actions) > 0 {
			return nil, errors.InvalidCharacter.SetDevMessage("role not presented in JWT")
		}

		return nil, nil
	}

	// Iterate over projects
	for project, metaData := range meta {
		isExistInProject := true
		// Get available actions by each role
		actionsMap, err := getRoleActions(*metaData.RoleID)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(actions); i++ {
			_, ok := actionsMap[actions[i]]
			if !ok {
				isExistInProject = false
			}
		}
		if isExistInProject {
			projects = append(projects, project)
		}
	}

	if len(projects) == 0 {
		return nil, errors.AccessDenied.SetDevMessage("don't have permissions to access")
	}

	return &projects, nil
}

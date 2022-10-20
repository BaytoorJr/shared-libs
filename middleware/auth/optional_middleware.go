package auth

import (
	"context"
	"github.com/BaytoorJr/shared-libs/errors"
	httpEncode "github.com/BaytoorJr/shared-libs/transport/http"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

// HTTPOptionalMiddleware Auth HTTP middleware with optional authentication
func HTTPOptionalMiddleware(h http.Handler, optionalActions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userActions string
		var hasOptional bool

		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			r.Header.Set("User-ID", "")
			r.Header.Set("User-Actions", "")
			r.Header.Set("Is-SuperAdmin", "false")
			r.Header.Set("Has-Optional-Actions", "false")
			h.ServeHTTP(w, r)
			return
		}

		token, err := jwt.Parse(tokenStr, getPublicKey)
		if err != nil {
			httpEncode.EncodeErrorResponse(
				context.Background(),
				errors.AccessDenied.SetDevMessage(err.Error()),
				w,
			)
			return
		}

		userData, err := getUserData(token)
		if err != nil {
			httpEncode.EncodeErrorResponse(
				context.Background(),
				errors.AccessDenied.SetDevMessage(err.Error()),
				w,
			)
			return
		}

		if len(optionalActions) > 0 {
			availableActions, err := checkUserOptionalPermissions(userData.RoleID, optionalActions)
			if err != nil {
				httpEncode.EncodeErrorResponse(
					context.Background(),
					errors.AccessDenied.SetDevMessage(err.Error()),
					w,
				)
				return
			}

			if len(availableActions) > 0 {
				hasOptional = true

				for i := 0; i < len(availableActions); i++ {
					userActions += availableActions[i]

					if i != len(availableActions)-1 {
						userActions += ","
					}
				}
			}
		}

		if userData.IsSuperAdmin {
			hasOptional = true
		}

		r.Header.Set("User-ID", userData.UserID)
		r.Header.Set("User-Mobile", userData.UserMobile)
		r.Header.Set("User-Actions", userActions)
		r.Header.Set("Is-SuperAdmin", strconv.FormatBool(userData.IsSuperAdmin))
		r.Header.Set("Has-Optional-Actions", strconv.FormatBool(hasOptional))

		if userData.ProfileID != nil {
			r.Header.Set("Profile-ID", *userData.ProfileID)
		}
		
		h.ServeHTTP(w, r)
	})
}

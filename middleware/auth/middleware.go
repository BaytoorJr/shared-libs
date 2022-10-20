package auth

import (
	"context"
	"github.com/BaytoorJr/shared-libs/errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"

	httpEncode "github.com/BaytoorJr/shared-libs/transport/http"
)

// HTTPMiddleware Auth HTTP middleware
func HTTPMiddleware(h http.Handler, actions, optionalActions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var optional bool

		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			httpEncode.EncodeErrorResponse(
				context.Background(),
				errors.AccessDenied.SetDevMessage("missing access token"),
				w,
			)
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

		if (len(actions) > 0 || len(optionalActions) > 0) && !userData.IsSuperAdmin {
			hasOptional, err := checkUserPermissions(userData.RoleID, actions, optionalActions)
			if err != nil {
				httpEncode.EncodeErrorResponse(
					context.Background(),
					errors.AccessDenied.SetDevMessage(err.Error()),
					w,
				)
				return
			}
			optional = hasOptional
		}

		if userData.IsSuperAdmin {
			optional = true
		}

		r.Header.Set("User-ID", userData.UserID)
		r.Header.Set("User-Mobile", userData.UserMobile)
		r.Header.Set("Is-SuperAdmin", strconv.FormatBool(userData.IsSuperAdmin))
		r.Header.Set("Has-Optional-Actions", strconv.FormatBool(optional))

		if userData.ProfileID != nil {
			r.Header.Set("Profile-ID", *userData.ProfileID)
		}
		
		h.ServeHTTP(w, r)
	})
}

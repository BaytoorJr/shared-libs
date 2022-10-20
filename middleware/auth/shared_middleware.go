package auth

import (
	"context"
	"github.com/BaytoorJr/shared-libs/errors"
	httpEncode "github.com/BaytoorJr/shared-libs/transport/http"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
)

// HTTPSharedMiddleware Auth HTTP middleware for shared services
func HTTPSharedMiddleware(h http.Handler, actions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var projects *[]string

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

		userData, err := getSharedUserData(token)
		if err != nil {
			httpEncode.EncodeErrorResponse(
				context.Background(),
				errors.AccessDenied.SetDevMessage(err.Error()),
				w,
			)
			return
		}

		if len(actions) > 0 {
			if !userData.IsSuperAdmin {
				projects, err = checkSharedUserPermissions(userData.Meta, actions)
				if err != nil {
					httpEncode.EncodeErrorResponse(
						context.Background(),
						errors.AccessDenied.SetDevMessage(err.Error()),
						w,
					)
					return
				}
			}
		}

		r.Header.Set("User-ID", userData.UserID)
		r.Header.Set("User-Mobile", userData.UserMobile)
		r.Header.Set("Is-SuperAdmin", strconv.FormatBool(userData.IsSuperAdmin))

		if projects != nil {
			r.Header.Set("Projects", strings.Join(*projects, ","))
		}

		if userData.ProfileID != nil {
			r.Header.Set("Profile-ID", *userData.ProfileID)
		}

		h.ServeHTTP(w, r)
	})
}

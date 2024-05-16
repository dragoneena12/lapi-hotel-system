package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dragoneena12/lapi-hotel-system/config"
	jwtauth "github.com/go-chi/jwtauth/v5"
)

type JWTAuthController struct{}

func NewJWTAuthController() *JWTAuthController {
	return &JWTAuthController{}
}

func (c *JWTAuthController) GetUserID(ctx context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get claims")
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get user ID")
	}
	return userID, nil
}

func JWTHasRole(ctx context.Context, requestedRole string) (bool, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get claims")
	}
	roles, ok := claims["roles"].([]string)
	if !ok {
		return false, fmt.Errorf("failed to get roles")
	}
	roles = append(roles, "USER")
	for _, r := range roles {
		if requestedRole == r {
			return true, nil
		}
	}
	return false, nil

}

func AuthMiddleware(cfg config.Config) func(http.Handler) http.Handler {
	tokenAuth := jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return jwtauth.Verifier(tokenAuth)
}

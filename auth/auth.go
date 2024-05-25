package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type (
	JWTAuthController struct{}
	userIDKey         struct{}
	roleKey           struct{}
)

var client = &http.Client{}

func NewJWTAuthController() *JWTAuthController {
	return &JWTAuthController{}
}

func (c *JWTAuthController) GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDKey{}).(string)
	if !ok {
		return "", fmt.Errorf("failed to get user ID from context")
	}
	return userID, nil
}

func JWTHasRole(ctx context.Context, requestedRole string) (bool, error) {
	roles, ok := ctx.Value(roleKey{}).([]string)
	if !ok {
		return false, fmt.Errorf("failed to get roles from context")
	}
	roles = append(roles, "USER")
	for _, r := range roles {
		if requestedRole == r {
			return true, nil
		}
	}
	return false, nil

}

func AuthMiddleWare(auth0Domain string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				next.ServeHTTP(w, r)
			}
			info, err := getUserInfo(token, auth0Domain)
			if err != nil {
				http.Error(w, "failed to get user info", http.StatusUnauthorized)
				slog.Error(fmt.Sprintf("failed to get user info: %s", err.Error()))
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey{}, info.UserID)
			ctx = context.WithValue(ctx, roleKey{}, info.Roles)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getUserInfo(token, auth0Domain string) (*Auth0UserInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/userinfo", auth0Domain), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	slog.Debug(fmt.Sprintf("request: %+v", req))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	slog.Debug(fmt.Sprintf("response body: %s", string(body)))
	info := &Auth0UserInfo{}
	if err := json.Unmarshal(body, info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return info, nil
}

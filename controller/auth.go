package controller

import "context"

type AuthController interface {
	GetUserID(ctx context.Context) (string, error)
}

type HasRole func(ctx context.Context, requestedRole string) (bool, error)

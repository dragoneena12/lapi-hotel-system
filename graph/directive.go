package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dragoneena12/lapi-hotel-system/controller"
	"github.com/dragoneena12/lapi-hotel-system/graph/model"
)

type directive struct {
	HasRole func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error)
}

func newDirective(hasRole controller.HasRole) *directive {
	return &directive{HasRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		ok, err := hasRole(ctx, role.String())
		if err != nil {
			return nil, fmt.Errorf("failed to get roles: %w", err)
		}
		if !ok {
			return nil, fmt.Errorf("role not allowed")
		}
		return next(ctx)
	}}
}

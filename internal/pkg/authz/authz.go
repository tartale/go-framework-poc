package authz

import (
	"context"
	"errors"

	"github.com/casbin/casbin"

	"github.com/logrhythm/go-framework-poc/internal/pkg/authn"
)

type AuthZ struct {
	enforcer *casbin.Enforcer
}

func NewAuthZ(modelPath string, policyPath string) (*AuthZ, error) {
	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}
	return &AuthZ{enforcer: e}, nil
}

func (o *AuthZ) Authorize(ctx context.Context, f func(params ...interface{}) error) error {
	jwt := ctx.Value("jwt").(authn.JWT)
	allowed, err := o.enforcer.Enforce(jwt.Sub, jwt.Obj, jwt.Act)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("Unauthorized")
	}
	return f(ctx)
}

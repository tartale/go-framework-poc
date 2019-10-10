package authz

import (
	"context"
	"errors"

	"github.com/casbin/casbin"

	"github.com/logrhythm/go-framework-poc/internal/pkg/authn"
)

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func NewAuthorizer(modelPath string, policyPath string) (*Authorizer, error) {
	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}
	return &Authorizer{enforcer: e}, nil
}

func (o *Authorizer) Authorize(ctx context.Context, fn ...interface{}) error {
	jwt := ctx.Value("jwt").(authn.JWT)
	allowed, err := o.enforcer.Enforce(jwt.Sub, jwt.Obj, jwt.Act)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("Unauthorized")
	}
	return fn(ctx)
}

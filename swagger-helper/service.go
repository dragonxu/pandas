// SPDX-License-Identifier: Apache-2.0

package swagger_helper

import (
	"context"

	"github.com/cloustone/pandas/pkg/errors"
	"github.com/mainflux/mainflux/logger"

	"github.com/cloustone/pandas/mainflux"
)

var (
	// ErrMalformedEntity indicates malformed entity specification (e.g.
	// invalid username or password).
	ErrMalformedEntity = errors.New("malformed entity specification")

	// ErrUnauthorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")

	// ErrConflict indicates that entity already exists.
	ErrConflict = errors.New("entity already exists")
)

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	// RetrieveDownstramSwagger retrieves data about downstream swagger with the provided
	// ID belonging to the model identified by the provided key.
	RetrieveDownstreamSwagger(context.Context, string, string) (DownstreamSwagger, error)
}

type swaggerService struct {
	auth     mainflux.AuthNServiceClient
	swaggers []DownstreamSwagger
}

var _ Service = (*swaggerService)(nil)

// New instantiates the swagger service implementation.
func New(auth mainflux.AuthNServiceClient, swaggerConfigs SwaggerHelperConfigs, logger logger.Logger) Service {
	return &swaggerService{
		auth:     auth,
		swaggers: swaggerConfigs.DownstreamSwaggers,
	}
}

func (ss *swaggerService) RetrieveDownstreamSwagger(ctx context.Context, token string, module string) (ds DownstreamSwagger, err error) {
	_, err = ss.auth.Identify(ctx, &mainflux.Token{Value: token})
	if err != nil {
		return DownstreamSwagger{}, ErrUnauthorizedAccess
	}
	for _, swagger := range ss.swaggers {
		if swagger.Name == module {
			return swagger, nil
		}
	}
	return DownstreamSwagger{}, ErrNotFound
}

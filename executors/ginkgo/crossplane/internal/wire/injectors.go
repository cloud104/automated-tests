//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
)

func SetUp(ctx context.Context, basename string) (*Test, func(), error) {
	wire.Build(providers)
	return nil, nil, nil
}

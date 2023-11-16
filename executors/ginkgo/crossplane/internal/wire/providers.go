package wire

import (
	"github.com/google/wire"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	wire.Struct(new(Test), "*"),
)

type Test struct {
}

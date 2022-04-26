package errors

import (
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

type ErrorMapper map[error]int
type ErrorMappers []ErrorMapper

var (
	Is           = errors.Is
	As           = errors.As
	New          = errors.New
	Unwrap       = errors.Unwrap
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

func NewErrorMappers() ErrorMappers {
	return ErrorMappers{
		userErrorMapper,
	}
}

var Module = fx.Options(
	fx.Provide(NewErrorMappers),
)

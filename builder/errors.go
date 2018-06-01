package builder

import (
	"fmt"
)

// ErrBuilderNotSupported builder not supported error
type ErrBuilderNotSupported struct {
	builderName string
}

func (err ErrBuilderNotSupported) Error() string {
	return fmt.Sprintf("%s builder not supported", err.builderName)
}

// NewErrBuilderNotSupported create ErrBuilderNotSupported instance
func NewErrBuilderNotSupported(builderName string) ErrBuilderNotSupported {
	return ErrBuilderNotSupported{
		builderName: builderName,
	}
}

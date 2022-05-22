package configor

import (
	"fmt"
	"strings"
)

// NotFound returns a not foung error on every lookup.
func NotFound() Resolver {
	return func(f *field) (string, error) {
		return "", fmt.Errorf("%w: path %s not found", ErrInvalidInput, strings.Join(f.path, "."))
	}
}

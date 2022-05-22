package configor

import (
	"os"
	"strings"
)

// Env resolves values from the environment.
func Env() Resolver {
	return func(f *field) (value string, err error) {
		var r []string

		for _, s := range f.path {
			s2 := strings.TrimSpace(strings.ToUpper(s))
			if s2 == "" {
				continue
			}

			r = append(r, s2)
		}

		return os.Getenv(strings.Join(r, "_")), nil
	}
}

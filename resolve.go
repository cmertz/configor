package configor

import (
	"strings"
)

// Resolve values for a configuration from passed Resolvers
func Resolve(config interface{}, resolvers ...Resolver) error {
	ls := leaves(config)
	for i := range ls {
		l := ls[i]
		for _, r := range resolvers {
			if l.isSet {
				break
			}

			v, err := r(&l)
			if err != nil {
				return err
			}

			ls[i] = l

			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}

			err = l.set(v)
			if err != nil {
				return err
			}

			ls[i] = l
		}
	}

	return nil
}

// prefix the paths of a list of fields with a path
// This is used when descending the config tree and
// building the paths.
func prefix(fields []field, prefix []string) []field {
	var r []field
	for _, f := range fields {
		r = append(r, field{
			t:    f.t,
			v:    f.v,
			path: append(prefix, f.path...),
		})
	}

	return r
}

// leaves returns all non-struct fields in the config
func leaves(config interface{}) []field {
	var res []field

	cur := []field{fieldFor(config)}

	for len(cur) > 0 {
		var next []field

		for _, c := range cur {
			if c.isStruct() {
				next = append(next, prefix(c.fields(), c.path)...)

				continue
			}

			res = append(res, c)
		}

		cur = next
	}

	return res
}

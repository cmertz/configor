package configor

// Prefix the config path.
func Prefix(prefix ...string) Resolver {
	return func(f *field) (string, error) {
		f.path = append(prefix, f.path...)

		return "", nil
	}
}

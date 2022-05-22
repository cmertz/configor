package configor

// Resolver is a resolver function for a value path
type Resolver func(f *field) (value string, err error)

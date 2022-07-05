package opts

// Opt is a type alias for a func that manipulates your opts.
type Opt[T any] func(*T)

// OptionsDefaulter is a factory interface for creating the default of your options.
type OptionsDefaulter[T any] interface {
	DefaultOptions() T
}

// DefaultApply uses the default factory to construct a default instance of the options and runs the Opt mutators on the
// result.
//
//     func GetUser(userId string, op ...opts.Opt[GetUserOpts]) (*User, error) {
//     	o := opts.DefaultApply(op...)
//     	...
//     }
//
//     type GetUserOpts struct {
//     	// Timeout if it takes longer than the specified time to get the user
//     	Timeout time.Duration
//     }
//
//     func (g GetUserOpts) DefaultOptions() GetUserOpts {
//     	return GetUserOpts{Timeout: 5 * time.Second}
//     }
//
//     func WithTimeout(t time.Duration) opts.Opt[GetUserOpts] {
//     	return func(o *GetUserOpts) {
//     		o.Timeout = t
//     	}
//     }
func DefaultApply[T OptionsDefaulter[T]](opts ...Opt[T]) T {
	a := (*new(T)).DefaultOptions()
	Apply[T](&a, opts...)
	return a
}

// Apply applies the mutator Opts to some instance of the options. See DefaultApply for a more in-depth example.
func Apply[T any](o *T, opts ...Opt[T]) {
	for _, v := range opts {
		v(o)
	}
}

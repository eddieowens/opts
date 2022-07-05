# opts

[![Go Report Card](https://goreportcard.com/badge/github.com/eddieowens/opts)](https://goreportcard.com/report/github.com/eddieowens/opts)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen.svg)](https://github.com/eddieowens/opts/blob/master/LICENSE)
[![godoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/eddieowens/opts?tab=doc)

Generic varargs options for your funcs

## Why

It's very common within several libraries to use vararg options e.g. in the
popular [grpc](https://pkg.go.dev/google.golang.org/grpc#WithAuthority) library

```go
grpc.Dial(":8080", grpc.WithBlock())
```

This provides a flexible and backwards compatible way to configure your APIs. The only problem is that it often requires
a good
deal of boilerplate. This library aims to help with reducing the boilerplate.

## Example

```go
package main

import (
	"github.com/eddieowens/opts"
	"time"
)

type GetUserOpts struct {
	// Timeout if it takes longer than the specified time to get the user
	Timeout time.Duration
}

// Implement (optional) opts.OptionsDefaulter interface to provide sane defaults.
func (g GetUserOpts) DefaultOptions() GetUserOpts {
	return GetUserOpts{Timeout: 5 * time.Second}
}

// Create option mutator func
func WithTimeout(t time.Duration) opts.Opt[GetUserOpts] {
	return func(o *GetUserOpts) {
		o.Timeout = t
	}
}

// Apply the options
func GetUser(userId string, op ...opts.Opt[GetUserOpts]) (*User, error) {
	o := opts.DefaultApply(op...)
	...
}

```

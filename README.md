# go-mod-di

Simple dependency injection container originally implemented by [michaelestrin](https://github.com/michaelestrin) for [edgexfoundry/go-mod-bootstrap](https://github.com/edgexfoundry/go-mod-bootstrap) in December 2019 and extracted here in June 2023 for general reuse.

Requires Go 1.18 or later.

## Usage

		package main

		import (
			"fmt"
			"github.com/michaelestrin/go-mod-di/pkg/di"
		)

		type foo struct {
			FooMessage string
		}

		func NewFoo(m string) *foo {
			return &foo{
				FooMessage: m,
			}
		}

		type bar struct {
			BarMessage string
			Foo        *foo
		}

		func NewBar(m string, foo *foo) *bar {
			return &bar{
				BarMessage: m,
				Foo:        foo,
			}
		}

		func main() {
			container := di.NewContainer(
				di.ServiceConstructorMap{
					"foo": func(get di.Get) interface{} {
						return NewFoo("fooMessage")
					},
					"bar": func(get di.Get) interface{} {
						return NewBar("barMessage", get("foo").(*foo))
					},
				})

			b := container.Get("bar").(*bar)
			fmt.Println(b.BarMessage)
			fmt.Println(b.Foo.FooMessage)
		}

### Background

[Dependency Injection in EdgeX Core Services](https://wiki.edgexfoundry.org/display/FA/Core+Working+Group?preview=%2F329472%2F37912747%2Fdi201910.pdf)
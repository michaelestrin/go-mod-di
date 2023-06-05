/**
 *  Copyright 2023 Michael Estrin.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"fmt"
	"go-mod-di/pkg/di"
)

type Foo struct {
	FooMessage string
}

func NewFoo(m string) *Foo {
	return &Foo{
		FooMessage: m,
	}
}

type Bar struct {
	BarMessage string
	Foo        *Foo
}

func NewBar(m string, foo *Foo) *Bar {
	return &Bar{
		BarMessage: m,
		Foo:        foo,
	}
}

func main() {
	container := di.NewContainer(
		di.ServiceConstructorMap{
			"Foo": func(get di.Get) any {
				return NewFoo("fooMessage")
			},
			"Bar": func(get di.Get) any {
				return NewBar("barMessage", get("Foo").(*Foo))
			},
		})

	b := container.Get("Bar").(*Bar)
	fmt.Println(b.BarMessage)
	fmt.Println(b.Foo.FooMessage)
}

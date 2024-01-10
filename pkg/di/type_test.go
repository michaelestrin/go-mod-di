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

package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeInstanceToNameReturnsExpectedNilTypeName(t *testing.T) {

	result := TypeInstanceToName(interface{}(nil))

	assert.Equal(t, "nil", result)
}

type s struct{}

func TestTypeInstanceToNameReturnsExpectedPackagePlusStructTypeName(t *testing.T) {

	result := TypeInstanceToName(s{})

	assert.Equal(t, "github.com/michaelestrin/go-mod-di/pkg/di.s", result)
}

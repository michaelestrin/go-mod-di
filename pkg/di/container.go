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

import "sync"

type Get func(serviceName string) any

// ServiceConstructor defines the contract for a function/closure to create a service.
type ServiceConstructor func(get Get) any

// ServiceConstructorMap maps a service name to a function/closure to create that service.
type ServiceConstructorMap map[string]ServiceConstructor

// service is an internal structure used to track a specific service's constructor and constructed instance.
type service struct {
	constructor ServiceConstructor
	instance    any
}

// Container is a receiver that maintains a list of services, their constructors, and their constructed instances in a
// thread-safe manner.
type Container struct {
	serviceMap map[string]service
	mutex      sync.RWMutex
}

// NewContainer is a factory method that returns an initialized Container receiver struct.
func NewContainer(serviceConstructors ServiceConstructorMap) *Container {
	c := Container{
		serviceMap: map[string]service{},
		mutex:      sync.RWMutex{},
	}
	if serviceConstructors != nil {
		c.Update(serviceConstructors)
	}
	return &c
}

// Update updates its internal serviceMap with the contents of the provided ServiceConstructorMap.
func (c *Container) Update(serviceConstructors ServiceConstructorMap) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for serviceName, constructor := range serviceConstructors {
		c.serviceMap[serviceName] = service{
			constructor: constructor,
			instance:    nil,
		}
	}
}

// get looks up the requested serviceName and, if it exists, returns a constructed instance.  If the requested service
// does not exist, it panics.  Get wraps instance construction in a singleton; the implementation assumes an instance,
// once constructed, will be reused and returned for all subsequent get(serviceName) calls.
func (c *Container) get(serviceName string) any {
	service, ok := c.serviceMap[serviceName]
	if !ok {
		panic("attempt to get unknown service \"" + serviceName + "\"")
	}
	if service.instance == nil {
		service.instance = service.constructor(c.get)
		c.serviceMap[serviceName] = service
	}
	return service.instance
}

// Get wraps get to make it thread-safe.
func (c *Container) Get(serviceName string) any {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.get(serviceName)
}

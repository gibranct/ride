package di

import (
	"fmt"
	"sync"
)

type objectMap map[string]any

type registry struct {
	objects objectMap
}

var instance *registry

var once sync.Once

func (r *registry) Add(key string, object any) {
	if r.containsKey(key) {
		return
	}

	r.objects[key] = object
}

func (r *registry) Get(key string) (object any, err error) {
	var ok bool

	if object, ok = r.objects[key]; !ok {
		return nil, fmt.Errorf("object was not registered for key: %s", key)
	}

	return object, nil
}

func (r *registry) containsKey(key string) bool {
	_, ok := r.objects[key]
	return ok
}

func (r *registry) Remove(key string) *registry {
	if _, ok := r.objects[key]; !ok {
		return r
	}

	delete(r.objects, key)

	return r
}

func NewRegistry() *registry {
	once.Do(func() {
		instance = &registry{
			objects: objectMap{},
		}
	})
	return instance
}

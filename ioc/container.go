package ioc

import (
	"reflect"
)

type Container struct {
	binds map[string]interface{}
	instances map[string]interface{}
}

func (c *Container) Bind(key string, concrete interface{})  {
	if reflect.TypeOf(concrete).Kind() == reflect.Func {
		if c.binds == nil {
			c.binds = make(map[string]interface{})
		}
		c.binds[key] = concrete
	} else {
		if c.instances == nil {
			c.instances = make(map[string]interface{})
		}
		c.instances[key] = concrete
	}
}

func (c *Container) Make(key string, params ...interface{})  interface{} {
	if obj, ok := c.instances[key]; ok {
		return obj
	}

	if fun, ok := c.binds[key]; ok {
		params = append([]interface{}{c}, params...)
		res := Apply(fun, params...)
		return res[0].Interface()
	}

	return nil
}

func Apply(f interface{}, params...interface{}) []reflect.Value {
	fun := reflect.ValueOf(f)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return fun.Call(in)
}


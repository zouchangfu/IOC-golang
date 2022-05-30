//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package config

import (
	"github.com/alibaba/ioc-golang/autowire"
	autowireconfig "github.com/alibaba/ioc-golang/extension/autowire/config"
)

func init() {
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigFloat64",
		Interface: new(ConfigFloat64),
		Factory: func() interface{} {
			return new(ConfigFloat64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigFloat64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configFloat64Interface)
			impl := i.(*ConfigFloat64)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigInt64",
		Interface: new(ConfigInt64),
		Factory: func() interface{} {
			return new(ConfigInt64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInt64Interface)
			impl := i.(*ConfigInt64)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigInt",
		Interface: new(ConfigInt),
		Factory: func() interface{} {
			return new(ConfigInt)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configIntInterface)
			impl := i.(*ConfigInt)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigMap",
		Interface: new(ConfigMap),
		Factory: func() interface{} {
			return new(ConfigMap)
		},
		ParamFactory: func() interface{} {
			return new(ConfigMap)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configMapInterface)
			impl := i.(*ConfigMap)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigSlice",
		Interface: new(ConfigSlice),
		Factory: func() interface{} {
			return new(ConfigSlice)
		},
		ParamFactory: func() interface{} {
			return new(ConfigSlice)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configSliceInterface)
			impl := i.(*ConfigSlice)
			return param.New(impl)
		},
	})
	autowireconfig.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias:     "ConfigString",
		Interface: new(ConfigString),
		Factory: func() interface{} {
			return new(ConfigString)
		},
		ParamFactory: func() interface{} {
			return new(ConfigString)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configStringInterface)
			impl := i.(*ConfigString)
			return param.New(impl)
		},
	})
}

type configFloat64Interface interface {
	New(impl *ConfigFloat64) (*ConfigFloat64, error)
}
type configInt64Interface interface {
	New(impl *ConfigInt64) (*ConfigInt64, error)
}
type configIntInterface interface {
	New(impl *ConfigInt) (*ConfigInt, error)
}
type configMapInterface interface {
	New(impl *ConfigMap) (*ConfigMap, error)
}
type configSliceInterface interface {
	New(impl *ConfigSlice) (*ConfigSlice, error)
}
type configStringInterface interface {
	New(impl *ConfigString) (*ConfigString, error)
}

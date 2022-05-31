//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package service2

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl1{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl2{}
		},
	})
}

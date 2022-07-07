//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package service2

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl1_{}
		},
	})
	impl1StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl1{}
		},
	}
	singleton.RegisterStructDescriptor(impl1StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl2_{}
		},
	})
	impl2StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl2{}
		},
	}
	singleton.RegisterStructDescriptor(impl2StructDescriptor)
}

type impl1_ struct {
	Hello_ func(name string) string
}

func (i *impl1_) Hello(name string) string {
	return i.Hello_(name)
}

type impl2_ struct {
	Hello_ func(name string) string
}

func (i *impl2_) Hello(name string) string {
	return i.Hello_(name)
}

type Impl1IOCInterface interface {
	Hello(name string) string
}

type Impl2IOCInterface interface {
	Hello(name string) string
}

func GetImpl1Singleton() (*Impl1, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Impl1)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl1)
	return impl, nil
}

func GetImpl1IOCInterfaceSingleton() (Impl1IOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl1)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(Impl1IOCInterface)
	return impl, nil
}

func GetImpl2Singleton() (*Impl2, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Impl2)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl2)
	return impl, nil
}

func GetImpl2IOCInterfaceSingleton() (Impl2IOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl2)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(Impl2IOCInterface)
	return impl, nil
}
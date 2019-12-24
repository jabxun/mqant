package consul

import ()
import "github.com/jabxun/mqant/registry"

func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}

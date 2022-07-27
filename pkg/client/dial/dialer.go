package dial

import (
	"k0s.io/pkg/client"
)

var _ client.Dialer = (*dialer)(nil)

func New(c client.Config) client.Dialer {
	return &dialer{
		c: c,
	}
}

type dialer struct {
	c client.Config
}

package hub

import (
	"github.com/btwiuse/invctrl/pkg/api"
)

type Slave struct {
	BidiStreamClient api.BidiStreamClient
	Info             string
}

package uuid

import (
	"github.com/google/uuid"
)

func New() string {
	randid := uuid.New()
	return randid.String()
}

func NewPet(s string) string {
	petid := uuid.NewSHA1(uuid.Nil, []byte(s))
	return petid.String()
}

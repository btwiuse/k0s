package asciitransport

import "k0s.io/k0s/pkg/asciiproto"

func O(b []byte) *asciiproto.Frame {
	return &asciiproto.Frame{Event: &asciiproto.Frame_O{O: b}}
}

func I(b []byte) *asciiproto.Frame {
	return &asciiproto.Frame{Event: &asciiproto.Frame_I{I: b}}
}

func R(w, h uint16) *asciiproto.Frame {
	return &asciiproto.Frame{Event: &asciiproto.Frame_R{R: uint32(w)<<16 | uint32(h)}}
}

func Uw(r uint32) uint16 {
	return uint16((r & 0xffff0000) >> 16)
}

func Uh(r uint32) uint16 {
	return uint16(r & 0xffff)
}

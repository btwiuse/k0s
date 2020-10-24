package asciitransport

func O(b []byte) *Frame {
        return &Frame{Event: &Frame_O{O: b}}
}

func I(b []byte) *Frame {
        return &Frame{Event: &Frame_I{I: b}}
}

func R(w, h uint16) *Frame {
        return &Frame{Event: &Frame_R{R: uint32(w)<<16 | uint32(h)}}
}

func Uw(r uint32) uint16 {
        return uint16((r & 0xffff0000) >> 16)
}

func Uh(r uint32) uint16 {
        return uint16(r & 0xffff)
}

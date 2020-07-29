package tunnel

import "time"

// since it's intended to be an http header, case doesn't matter
const FingerPrintHeader = "X-Backend-ID"
const KeepaliveInterval = 6 * time.Second

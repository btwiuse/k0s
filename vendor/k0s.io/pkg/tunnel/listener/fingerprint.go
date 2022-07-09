package listener

import (
	"strconv"
	"time"
)

var fingerprint string = strconv.FormatInt(time.Now().UnixNano(), 10)

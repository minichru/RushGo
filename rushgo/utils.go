package rushgo

import (
	"time"
)


func Second(seconds int) time.Duration {
    return time.Duration(seconds) * time.Second
}


package snowflake

import (
	"sync/atomic"
	"time"
)

type IDGen struct {
	id int64 // encodes last ts + last seq
}

func NewSnowflakeGen() *IDGen {
	return &IDGen{}
}

func (g *IDGen) NewID() int64 {
	for {
		ts := time.Now().UnixMilli()

		candidateID := atomic.AddInt64(&g.id, 1)
		if ts == candidateID>>16 {
			// Got new incremented ID during the same ms, great, we're done!
			return candidateID
		}

		// Timestamp has changed since last generation.
		// We must reset the ID with a new ts and seq=0
		// If no other thread touched g.id, it should be
		// equals to candidateID returned by AddInt64.

		if atomic.CompareAndSwapInt64(&g.id, candidateID, ts<<16) {
			// Reset successful, we can return the new ID
			// There is no need to OR anything, because seq=0
			return ts << 16
		}
		// Not successful... It seems that another thread touched g.id
		// before we were able to update it. We need to retry.
	}
}

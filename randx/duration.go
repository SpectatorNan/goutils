package tools

import (
	"github.com/zeromicro/go-zero/core/mathx"
	"math/rand"
	"time"
)

// RandDuration is a struct that provides methods to generate random durations
type RandDuration struct {
	deviation          float64
	unstableExpiryTime mathx.Unstable
}

func NewRandDuration(deviation float64) *RandDuration {
	return &RandDuration{
		deviation:          deviation,
		unstableExpiryTime: mathx.NewUnstable(deviation),
	}
}

// AroundDuration generates a random duration around a base duration
// range: [(1 - deviation) * base, (1 + deviation) * base]
func (r *RandDuration) AroundDuration(base time.Duration) time.Duration {
	return r.unstableExpiryTime.AroundDuration(base)
}

// JitterDuration generates a random duration with jitter
// range: [base, base * (1 + deviation)]
func (r *RandDuration) JitterDuration(base time.Duration) time.Duration {
	jitter := time.Duration(rand.Float64() * float64(base) * r.deviation)
	return base + jitter
}

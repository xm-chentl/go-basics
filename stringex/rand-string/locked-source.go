package randstring

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	Src rand.Source64

	lk sync.Mutex
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.Src.Int63()
	r.lk.Unlock()
	return
}

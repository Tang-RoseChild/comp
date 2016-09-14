package midware

import (
	"net/http"
	"sync"
)

type HandlerLimit struct {
	Mutex      *sync.RWMutex
	Max        int
	LimitIpMap map[string]chan struct{}
}

// LimitHandler default max req/ip is 10
func LimitHandler(limit *HandlerLimit, handler http.HandlerFunc) http.HandlerFunc {
	switch {
	case limit == nil:
		limit = &HandlerLimit{
			Mutex:      &sync.RWMutex{},
			Max:        10,
			LimitIpMap: make(map[string]chan struct{}),
		}
	case limit.Mutex == nil:
		limit.Mutex = &sync.RWMutex{}
		fallthrough
	case limit.Max == 0:
		limit.Max = 10
		fallthrough
	case limit.LimitIpMap == nil:
		limit.LimitIpMap = make(map[string]chan struct{})

	}

	return func(w http.ResponseWriter, r *http.Request) {

		if _, ok := limit.LimitIpMap[r.Host]; !ok {
			limit.Mutex.Lock()
			limit.LimitIpMap[r.Host] = make(chan struct{}, limit.Max)
			limit.Mutex.Unlock()
		}
		limit.LimitIpMap[r.Host] <- struct{}{}
		defer func() {
			<-limit.LimitIpMap[r.Host]
		}()

		handler(w, r)
		return
	}

}

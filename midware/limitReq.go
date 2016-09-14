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

func LimitHandler(limit *HandlerLimit, handler http.HandlerFunc) http.HandlerFunc {

	if limit.Mutex == nil {
		limit.Mutex = &sync.RWMutex{}
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

package midware

import "net/http"

// Wrapper wrap a http.HandlerFunc for decorating
// param is for wrapping,
// return val is for doing what you want to wrap,like adding a log
type Wrapper func(http.HandlerFunc) http.HandlerFunc

// Wrap could use different wraps to decorate a http.HandlerFunc,
// like wanting a log,outh check,params check and so on
func Wrap(h http.HandlerFunc, wrappers ...Wrapper) http.HandlerFunc {
	// reverse to wrap,more readable,like Wrap(index,log(x),auth(x)...)

	for i := len(wrappers) - 1; i >= 0; i-- {

		h = wrappers[i](h)
	}
	return h
}

// OriginChecker origin check middleware
// if OriginCheck is ok, will return true
// otherwise return false, OriginNotPass using to handle it
type OriginChecker interface {
	OriginCheck(http.ResponseWriter, *http.Request) bool
	OriginNotPass()
}

// CheckOrigin use the OriginChecker to check and return a Wrapper
func CheckOrigin(checker OriginChecker) Wrapper {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if checker.OriginCheck(w, r) {
				h(w, r)
				return
			}
			checker.OriginNotPass()
		}
	}
}

// Tokener token check middleware
// TokenCheck need to return true if ok
// if not ok, TokenNotPass use to handle it
type Tokener interface {
	TokenCheck(http.ResponseWriter, *http.Request) bool
	TokenNotPass()
}

// CheckOrigin use the Tokener to check and return a Wrapper
func CheckToken(tokener Tokener) Wrapper {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if tokener.TokenCheck(w, r) {
				h(w, r)
				return
			}
			tokener.TokenNotPass()
		}
	}
}

/*

old ones

// // WithOrigin use OriginChecker to check. if check ok,h will take over,
// // otherwise call OriginNotPass to handle the err
// func WithOrigin(checker OriginChecker, h http.HandlerFunc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if checker.OriginCheck(w, r) {
// 			h(w, r)
// 			return
// 		}
// 		// TODO: not check in handle
// 		checker.OriginNotPass()

// 	}
// }



// // WithToken use Tokener to check. if check ok,h will take over,
// // otherwise use the TokenNotPass to handle err.
// func WithToken(tokener Tokener, h http.HandlerFunc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if tokener.TokenCheck(w, r) {
// 			h(w, r)
// 			return
// 		}
// 		// TODO: not check in handle
// 		tokener.TokenNotPass()

// 	}
// }

*/

package midware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWrap(t *testing.T) {
	// http.HandleFunc("/test", Wrap(testIndex, CheckOrigin(testOrigin{}), CheckToken(testTokener{})))
	// http.ListenAndServe(":8081", nil)

	Wrap(testIndex, wraperA(), wraperXX())(nil, nil)
}

func wraperA() Wrapper {
	return func(h http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("in wraperA")
			h(w, r)
		}
	}
}

func wraperXX() Wrapper {
	return func(h http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("in wraperXX")
			h(w, r)
		}
	}
}

func testIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in testIndex")
}

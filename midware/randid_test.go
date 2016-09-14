package midware

import (
	"fmt"
	"testing"
)

var ids = make(map[string]struct{})

func BenchmarkRandID6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id, _ := GenerateRandID(6)
		fmt.Println("i ", i, " id ", id)
		if _, ok := ids[id]; ok {
			panic(fmt.Sprintln(id, " id exist "))
		} else {
			ids[id] = struct{}{}
		}

	}
	fmt.Println("in benching")
}

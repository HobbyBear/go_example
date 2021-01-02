package pool

import (
	"fmt"
	"testing"
	"time"
)
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	//GiB // 1073741824
	//TiB // 1099511627776             (超过了int32的范围)
	//PiB // 1125899906842624
	//EiB // 1152921504606846976
	//ZiB // 1180591620717411303424    (超过了int64的范围)
	//YiB // 1208925819614629174706176
)

var curMem uint64
func BenchmarkPool_Submit(t *testing.B) {
	p, err := NewPool()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < t.N; i++ {
		num := i
		p.Submit(func() {
			time.Sleep(10*time.Millisecond)
			fmt.Println(num)
		})
	}
}

func Test01(t *testing.T) {
	arr := []int{0,1,2,3}
	fmt.Println(arr[0:2])
}
package size

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"gopkg.in/mgo.v2/bson"
)

type StudentAnswerCache struct {
	StudentID  string    `msgpack:"student_id"`
	TitleID    int64     `msgpack:"title_id"`
	Answer     string    `msgpack:"answer"`
	SubmitTime time.Time `msgpack:"submit_time"`
	Duration   int64     `msgpack:"duration"`
	Status     int       `msgpack:"status"`
}

func TestSize(t *testing.T) {
	s := bson.NewObjectId()
	var s2 StudentAnswerCache

	//m := map[string]StudentAnswerCache{"5eb51da252a500aedad33a48": s}
	//m2 := map[string]StudentAnswerCache{"5eb51da252a500aedad33a48": s2}
	//fmt.Println(SizeStruct(m))
	//fmt.Println(SizeStruct(m2))
	fmt.Println(SizeStruct(s))
	fmt.Println(SizeStruct(s2))
	var num int64
	nums := struct {
		Num  int64  `json:"num"`
		Name string `json:"name"`
		Num2 int64  `json:"num_2"`
	}{Num: num, Num2: num, Name: ""}
	fmt.Println(SizeStruct(nums))
}

func SizeStruct(data interface{}) int {
	return sizeof(reflect.ValueOf(data))
}

func sizeof(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Map:
		sum := 0
		keys := v.MapKeys()
		for i := 0; i < len(keys); i++ {
			mapkey := keys[i]
			s := sizeof(mapkey)
			if s < 0 {
				return -1
			}
			sum += s
			s = sizeof(v.MapIndex(mapkey))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum
	case reflect.Slice, reflect.Array:
		sum := 0
		for i, n := 0, v.Len(); i < n; i++ {
			s := sizeof(v.Index(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.String:
		sum := 0
		for i, n := 0, v.Len(); i < n; i++ {
			s := sizeof(v.Index(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Ptr, reflect.Interface:
		p := (*[]byte)(unsafe.Pointer(v.Pointer()))
		if p == nil {
			return 0
		}
		return sizeof(v.Elem())
	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := sizeof(v.Field(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Int:
		return int(v.Type().Size())

	default:
		fmt.Println("t.Kind() no found:", v.Kind())
	}

	return -1
}

func Test02(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 3
	for c := range ch {
		fmt.Println(c)
	}
}

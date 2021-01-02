package _interface

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *student) string() {
	fmt.Println("student")
}

func Test01(t *testing.T) {
	var s student
	fmt.Println(s)
}

func Test02(t *testing.T) {
	s := new(student)
	fmt.Println(s)
}

func Test03(t *testing.T) {
	var arr []string
	fmt.Println(arr)
}

func Test04(t *testing.T) {
	var m map[string]int
	// map 取不到会取对应value值得默认值
	name, ok := m["name"]
	fmt.Println(name, ok)
	fmt.Println(m["name"])
}

func Test05(t *testing.T) {
	var s *student
	fmt.Println(s)
}

type person interface {
	say() string
	hello() string
}

type stu struct {
	Name string `json:"name"`
}

func (s *stu) say() string {
	s.Name = "wky"
	fmt.Println(s.Name)
	return s.Name
}

func (s *stu) hello() string {
	fmt.Println(s.Name)
	return s.Name
}

func Test06(t *testing.T) {
	// 接口类型转换为对象类型只能用类型断言

	// 对象类型转换，对象转接口用强制类型转换
	s := (*stu)(nil)
	fmt.Println(reflect.ValueOf(s))
}

func Test07(t *testing.T) {
	s := new(stu)
	fmt.Println(s.Name)
}

func Test08(t *testing.T) {
	type student struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Class int `json:"class"`
	}
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := person{
		Name: "xch",
		Age:  12,
	}
	byte, _ := json.Marshal(p)
	var stu student
	err := json.Unmarshal(byte, &stu)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stu)
}

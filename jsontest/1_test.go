package jsontest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type Num int64

func (p *Num) UnmarshalJSON(bytes []byte) error {
	i, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *Num) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func Test01(t *testing.T) {
	var n Num
	s := "132"
	err := json.Unmarshal([]byte(s), &n)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n)
}

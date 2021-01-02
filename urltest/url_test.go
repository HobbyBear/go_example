package urltest

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func Test01(t *testing.T) {
	var people struct {
		Ids []int64 `url:"ids"`
	}

	people.Ids = []int64{1, 2, 3}

	val, err := query.Values(people)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val.Encode())
}

func Test02(t *testing.T) {
	s, err := primitive.ObjectIDFromHex("000000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s.IsZero())
	fmt.Println(s)
}

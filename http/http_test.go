package main

import (
	"net/http"
	"testing"
)

func Test01(t *testing.T) {
	_,err := http.Get("http://localhost:9091")
	if err != nil{
		t.Fatal(err)
	}

}

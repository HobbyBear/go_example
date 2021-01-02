package main

import "fmt"

type Handler interface {
	Do() error
	SetNext(h Handler) Handler
	Run() error
}

type DefaultHandler struct {
	Next Handler
}

func (d *DefaultHandler) Do() error {
	fmt.Println("默认执行器")
	return nil
}

// 组合的方式复用setNext和Run方法
func (d *DefaultHandler) SetNext(h Handler) Handler {
	d.Next = h
	return h
}

func (d *DefaultHandler) Run() error {
	if d.Next != nil {
		d.Next.Do()
		return d.Next.Run()
	}
	return nil
}

type NullHandler struct {
	DefaultHandler
}

func (p *NullHandler) Do() error {
	return nil
}

type ParamValidHandler struct {
	DefaultHandler
}

func (p *ParamValidHandler) Do() error {
	fmt.Println("参数校验开始")
	return nil
}

type Param2ValidHandler struct {
	DefaultHandler
}

func (p *Param2ValidHandler) Do() error {
	fmt.Println("参数校验2开始")
	return nil
}

func main() {
	var n NullHandler
	var p ParamValidHandler
	var p2 Param2ValidHandler
	n.SetNext(&p).SetNext(&p2)
	n.Run()
}

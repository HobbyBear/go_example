package zaptest

import (
	"fmt"
	"testing"

	"app/msgbus/logtest"

	"go.uber.org/zap"
)

func Test03(t *testing.T) {
	logtest.Logger.Info("test", zap.String("name", "xch"))
	arr := []int{1, 2}
	fmt.Println(arr)
	arr1 := append(arr, 3)
	fmt.Println(arr)
	fmt.Println(arr1)
}

func Test02(t *testing.T) {

	nullHandler := &NullHandler{}
	nullHandler.SetNext(&ArgumentsHandler{})

	// 开始执行业务
	if err := nullHandler.Run(&Context{}); err != nil {
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	fmt.Println("success")
	return
}

type TestHandler struct {
}

func (t *TestHandler) Do(c *Context) error {
	fmt.Println("=================")
	return nil
}

func (t *TestHandler) SetNext(h Handler) Handler {
	return nil
}

func (t *TestHandler) Run(c *Context) error {
	return nil
}

type Context struct {
}

type Handler interface {
	Do(c *Context) error
	SetNext(h Handler) Handler
	Run(c *Context) error
}

type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

func (n *Next) Run(c *Context) (err error) {
	if n.nextHandler != nil {
		if err = (n.nextHandler).Do(c); err != nil {
			return
		}
		return (n.nextHandler).Run(c)
	}
	return
}

type NullHandler struct {
	Next
}

func (h *NullHandler) Do(c *Context) (err error) {
	return
}

type ArgumentsHandler struct {
	Next
}

func (h *ArgumentsHandler) Do(c *Context) (err error) {
	fmt.Println("参数校验成功")
	return
}



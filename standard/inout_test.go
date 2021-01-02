package standard

import (
	"container/list"
	"fmt"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestIn(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "xch"
	m["haha"] = "h1"
	type people struct {
		Name string `json:"name"`
	}
	p := new(people)
	p.Name = "xch"
	arr := []*people{p}
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"name": arr,
	}).Info("success", "fail")
}

func TestList(t *testing.T) {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestAddStructToList(t *testing.T) {
	type people struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	list := list.New()
	list.PushBack(people{
		Name: "xch",
		Age:  10,
	})
	for e := list.Front(); e != nil; e = e.Next() {

		fmt.Println()
	}
}

func Test2(t *testing.T) {
	arr := [3][3]int{
		{0, 1, 2}, /*  第一行索引为 0 */
		{4, 5, 6}, /*  第二行索引为 1 */
		{8, 9, 10}}
	Upserver(arr)

}

//使用交换值实现

func Upserver(arr [3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	fmt.Println(arr)
}

func TestScan(t *testing.T) {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 3
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		wg.Done()
		num := i
		go func() {
			wg.Wait()
			fmt.Println(m[num])
		}()
	}
}

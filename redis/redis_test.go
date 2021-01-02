package redis

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/xiaojiaoyu100/cast"
)

const script = `
	local key1 = KEYS[1]
	local key2 = KEYS[2]
	local val1 = ARGV[1]
	return redis.call('set',key1,val1)
	

`

var limitAnswerKindScripts = redis.NewScript(`
	local answerSetKey = tostring(KEYS[1])		
	local answer = tostring(ARGV[1])
	local length = redis.call("scard",answerSetKey)
	local expiration = tonumber(ARGV[2])
if length >= 1 then
	return 0
end
	redis.call("sadd",answerSetKey,answer)
	redis.call("expire",answerSetKey,expiration)
	return 1
`)

type customerHook struct {
}

func (c customerHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	fmt.Println(cmd.Args())
	return ctx, nil
}

func (c customerHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	fmt.Println(cmd.Args())
	return nil
}

func (c customerHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return nil, nil
}

func (c customerHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}

func Test01(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	//res, err := client.ZRevRank("12","5").Result()
	//if err != nil{
	//	t.Fatal(err)
	//}
	//fmt.Println(res)
	var c customerHook
	client.AddHook(c)
	a, err := client.Del("*").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func Test02(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	redis.NewRing(&redis.RingOptions{
		Addrs: nil,
	})

	pipeline := client.Pipeline()
	pipeline.HGet("name2", "xch")
	cmdList, err := pipeline.Exec()
	if err != nil {
		t.Fatal(err)
	}
	for _, cmd := range cmdList {
		fmt.Println(cmd.Args())
		fmt.Println(cmd.Name())
		fmt.Println(cmd.Err())
	}
}

func Test3(t *testing.T) {
	c, _ := cast.New()
	for i := 0; i <= 100; i++ {
		req := c.NewRequest().WithPath("https://title-cdn.xhwx100.com/dev/teaching/5dd9f56bd26176000101d7b9.json").Get()
		resp, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if !resp.Success() {
			t.Fatal(resp.StatusCode())
		}
		fmt.Println(string(resp.Body()))
	}

}

func Test4(t *testing.T) {

	for i := 0; i < 100; i++ {
		resp, err := http.Get("https://title-cdn.xhwx100.com/dev/teaching/5dd9f56bd26176000101d7b9.json")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatal(resp.StatusCode)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}

}

func Test5(t *testing.T) {
	client := &http.Client{}
	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest(http.MethodGet, "https://title-cdn.xhwx100.com/dev/teaching/5dd9f56bd26176000101d7b9.json", nil)
		resp, _ := client.Do(req)
		if resp.StatusCode != 200 {
			t.Fatal(resp.StatusCode)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}

}

func Test07(t *testing.T) {

}

func reqBatch(idList []int) {
	var wg sync.WaitGroup
	wg.Add(len(idList))

	// 循环里用协程池去请求
	for i := 1; i <= len(idList); i++ {
		go func() {
			defer wg.Done()
			// do something
		}()
	}
	wg.Wait()
}

func Test06(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	_, err := client.Set("person", []byte("hahah"), 24*time.Hour).Result()
	if err != nil {
		t.Fatal(err)
	}

}

func Test7(t *testing.T) {
	type people struct {
		Name string `json:"name"`
	}

	testM := make(map[int]*people)
	testM[1] = &people{Name: "xch"}
	testM[2] = &people{Name: "xch2"}
	var l sync.RWMutex
	go func() {
		var num int
		for {
			l.Lock()
			num++
			testM[num] = &people{Name: "xch3"}
			fmt.Println(testM[num])
			l.Unlock()
		}
	}()

	go func() {
		var num int
		for {
			l.RLock()
			p := testM[num]
			if p != nil {
				p.Name = "xch3"
			}
			l.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)

}

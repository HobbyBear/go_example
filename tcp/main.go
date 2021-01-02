package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal(err)
	}

	//err = conn.SetWriteDeadline(time.Now().Add(time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}


	// 调用返回的连接对象提供的 Write 方法发送请求
	_, err = conn.Write([]byte("GET / HTTP/1.1\nUser-Agent: PostmanRuntime/7.26.5\nAccept: */*\nCache-Control: no-cache\nPostman-Token: 44189221-729b-4b13-878d-ee1651d96c86\nHost: localhost:9091\nAccept-Encoding: gzip, deflate, br\nConnection: keep-alive"))

	if err != nil{
		log.Println(err,"write fail")
		return
	}

	// 通过连接对象提供的 Read 方法读取所有响应数据
	result, err := readFully(conn)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(result))

}

func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应数据后主动关闭连接
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [50]byte
	for {
		n, err := conn.Read(buf[0:])
		time.Sleep(12 * time.Millisecond)
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
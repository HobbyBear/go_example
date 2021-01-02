package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type NodeInfo struct {
	VendorCode  string `json:"vendorCode"`  // 商家编码
	VendorName  string `json:"vendorName"`  // 商家名称
	WaybillCode string `json:"waybillCode"` // 运单号
	TraceNode   string `json:"traceNode"`   // 操作节点
	TraceMark   string `json:"traceMark"`   // 操作描述
	OrderId     string `json:"orderId"`     // 订单号  ECLP单号 仓配的都是ESL单号
	Operator    string `json:"operator"`    // 操作员
	OperateTime string `json:"operateTime"` // 操作时间
}

func main() {
	urlPath := "https://preop.xhwx100.com/gorilla/api/v2.0/textbook/send_packet/delivery_info/callback"
	data := make(url.Values)
	var node = NodeInfo{
		VendorCode:  "021K10000",
		VendorName:  "XXXX有限公司",
		WaybillCode: "JDVC05101897701",
		TraceNode:   "逆向发货",
		TraceMark:   "逆向发货",
		OrderId:     "ESL74788228357651",
		Operator:    "张三",
		OperateTime: "2017-12-15 20:25:20",
	}
	nodeInfo,err  := json.Marshal(node)
	if err != nil{
		fmt.Println(err)
		return
	}

	requestBody := string(nodeInfo)
	timeStamp := "2019-5-12 00:00:00"
	data["message_id"] = []string{"123"}
	data["format"] = []string{"123"}
	data["request_body"] = []string{requestBody}
	data["timestamp"] = []string{timeStamp}
	h := md5.New()
	h.Write([]byte(requestBody + timeStamp + "xhwx100jc12345687"))
	myToken := hex.EncodeToString(h.Sum(nil))

	data["token"] = []string{myToken}
	res,err := http.PostForm(urlPath, data)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	respData,err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(respData))
}

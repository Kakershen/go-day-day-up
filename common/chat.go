// 聊天机器人
package main

/*
func main() {
	var input string
	fmt.Println("Entry EOF to shutdown:")
	channel := make(chan string)
	// 结束世关闭通道
	defer close(channel)
	// 启动goroutine运行机器人回答线程
	go robot(channel, string(rand.Int63()))
	for {
		// 从命令行读取输入
		fmt.Scanf("%s", &input)
		channel <- input
		if input == "EOF" {
			fmt.Println("Bye!")
			break
		}
	}

}
*/


import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 请求结构体
type requestBody struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userid"`
}

// 相应结构体
type resonseBody struct {
	Code int      `json:"code"`
	Text string   `json:"text"`
	List []string `json:"list"`
	Url  string   `json:"url"`
}

// 请求机器人
func robot(inputChan <-chan string, userid string) {
	for {
		// 不停的监听通道输入
		input := <-inputChan
		if input == "EOF" {
			break
		}
		// 构造请求体
		reqData := &requestBody{
			Key:    "792bcf45156d488c92e9d11da494b085",
			Info:   input,
			UserId: userid,
		}
		byteData, _ := json.Marshal(&reqData)
		//请求聊天机器人接口
		req, err := http.NewRequest("POST", "http://www.tuling123.com/openapi/api",
			bytes.NewReader(byteData))
		req.Header.Set("Content-Type", "application/json;chartset=UTF-8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Network err!", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			var respData resonseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI:" + respData.Text)
		}
		if resp != nil {
			resp.Body.Close()
		}
	}
}

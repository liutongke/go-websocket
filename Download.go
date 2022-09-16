package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := ""
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "insert-always=202489098.36895.0000")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, _ := os.OpenFile("D:/phpstudy/project/private/go_work/go-websocket/test.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(body)) //将数据先写入缓存
	writer.Flush()
	fmt.Println(string(body))
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type GetNameResponse struct {
	Code uint   `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func ApiSdk(apiUrl, name string) (GNR GetNameResponse) {
	data := url.Values{}
	data.Set("username", name)
	// 解析完整url
	uri, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return
	}
	uri.RawQuery = data.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// 在请求头中添加自定义的 key
	req.Header.Set("Key", "test")

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 通用写法
	// var data map[string]interface{}
	//    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	//        return nil, err
	//    }

	err = json.Unmarshal(b, &GNR)
	if err != nil {
		return
	}

	fmt.Println(GNR.Msg)
	return
}

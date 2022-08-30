package main

import (
	"dictionary-go/model"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var (
		from  string
		to    string
		query string
	)
	// go run .\main.go --from auto --to en --q 你好
	flag.StringVar(&from, "from", "auto", "从那个语言翻译过来")
	flag.StringVar(&to, "to", "en", "翻译语言")
	flag.StringVar(&query, "q", "", "翻译的文本")
	flag.Parse()
	tran := model.NewTranslateModeler(query, from, to)
	values := tran.ToValues()
	resp, err := http.PostForm(model.Url, values) //调用一个post请求 并且上传表单
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	res := model.TransRespone{}            //初始化一个结构体接收返回消息
	body, err := ioutil.ReadAll(resp.Body) //读取响应的响应体
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &res)    //反序列化到struct中
	for _, v := range res.Trans { //输出源文本和目标文本
		fmt.Println("原文本", v.Src)
		fmt.Println("翻译文本", v.Dst)
	}

}

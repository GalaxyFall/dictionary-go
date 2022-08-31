package model

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

//申请的信息  iD和密钥
const (
	UerAppID    = "20220816001309619"
	UerPassword = "W_h0op7WCB9qcLUVl34B"
)

//百度翻译api接口
var Url = "http://api.fanyi.baidu.com/api/trans/vip/translate"

//抽象成一个model
type TranslateModel struct {
	Q     string
	From  string
	To    string
	Appid string
	Salt  int
	Sign  string
}

//接收返回json的结构体  拆分
type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type TransRespone struct {
	From  string        `json:"from"`
	To    string        `json:"to"`
	Trans []TransResult `json:"trans_result"`
}

//初始化一个翻译模块
func NewTranslateModeler(q, from, to string) TranslateModel {
	tran := TranslateModel{
		Q:    q,
		From: from,
		To:   to,
	}
	tran.Appid = UerAppID
	tran.Salt = time.Now().Second()
	content := UerAppID + q + strconv.Itoa(tran.Salt) + UerPassword
	sign := SumString(content) //计算sign值
	tran.Sign = sign
	return tran
}

//将对应的请求表单组装好
func (tran TranslateModel) ToValues() url.Values {
	values := url.Values{
		"q":     {tran.Q},
		"from":  {tran.From},
		"to":    {tran.To},
		"appid": {tran.Appid},
		"salt":  {strconv.Itoa(tran.Salt)},
		"sign":  {tran.Sign},
	}
	return values
}

//计算文本的md5值
func SumString(content string) string {
	bys := md5.Sum([]byte(content))
	//bys := md5.Sum([]byte(content))//这个md5.Sum返回的是数组,不是切片哦
	value := hex.EncodeToString(bys[0:]) //根据数组构造一个切片
	return value
}

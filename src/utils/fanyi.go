package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type query struct {
	Q     string `json:"q"`
	From  string `json:"from"`
	To    string `json:"to"`
	Appid string `json:"appid"`
	Salt  string `json:"salt"`
	Sign  string `json:"sign"`
}

type transResult struct {
	Src string
	Dst string
}

type res struct {
	From        string
	To          string
	TransResult []transResult `json:"trans_result"`
	ErrorCode   int           `json:"error_code"`
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func FanYi(q string) string {
	rand.Seed(time.Now().UnixNano())
	salt := strconv.Itoa(int(time.Now().Unix()))
	appid := "20220819001313212"
	f := &query{
		Q:     q,
		From:  "auto",
		To:    "zh",
		Appid: appid,
		Salt:  salt,
		Sign:  MD5(appid + q + salt + "csm4w2Pc9FhiGUL3oUgl"),
	}
	url := "http://fanyi-api.baidu.com/api/trans/vip/translate" + StructToParm(f)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var res res
	json.Unmarshal(body, &res)
	return res.TransResult[0].Dst
}

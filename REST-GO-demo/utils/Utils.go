package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//使用sha512加密
func Sha(s string, secret string) string {
	hash := sha512.New()
	hash.Write([]byte(s + secret))
	return hex.EncodeToString(hash.Sum(nil))
}

// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// isSign: 是否需要签名
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]interface{}, isSign bool) string {
	var strParams string
	//签名加入随机数
	if isSign {
		mapParams["nonce"] = time.Now().UnixNano() / 1000
	}

	for key, value := range mapParams {
		strParams += (key + "=" + fmt.Sprint(value) + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// Http Get请求基础函数, 通过封装Go语言Http请求
// strUrl: 请求的URL
// strParams: string类型的请求参数, user=lxz&pwd=lxz
// hostURL服务器主机
// secretKey加密秘钥
// return: 请求结果
func HttpGetRequest(strUrl string, mapParams map[string]interface{}, hostURL, secretKey string) string {
	httpClient := &http.Client{}

	//判断参数是否含有key，有key需要签名
	isSign := false
	if _, ok := mapParams["key"]; ok {
		isSign = true
	}

	strRequestUrl := ""
	//拼接请求方法和参数
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams, isSign)
		strRequestUrl = strUrl + "?" + strParams
	}

	//签名并添加签名参数
	if isSign {
		sign := Sha(strRequestUrl, secretKey)
		strRequestUrl = strRequestUrl + "&sign=" + sign
	}

	//拼接完整请求地址
	payLoad := hostURL + strRequestUrl
	// log.Println(payLoad)
	// 构建Request, 并且按官方要求添加Http Header
	request, err := http.NewRequest("GET", payLoad, nil)
	if nil != err {
		return err.Error()
	}

	// 发出请求
	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()
	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

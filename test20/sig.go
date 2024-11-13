package main

import (
	"crypto/md5"
	"encoding/hex"

	"net/url"
	"sort"
	"strings"
	"time"
)

// 初始化安全部签名SDK (模拟实现)
func InitWsgsigSign(bizId, version, osType string) {
	// 初始化签名SDK，这里仅是个示例
	// 实际实现可能会涉及到SDK调用或环境配置
}

// 序列化参数
func Serialize(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, key := range keys {
		val := params[key]
		valStr := val
		parts = append(parts, url.QueryEscape(key)+"="+url.QueryEscape(valStr))
	}
	return strings.Join(parts, "&")
}

// 生成MD5签名
func Sig(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// 过滤无效参数
func FilterUndefinedParams(params map[string]string) map[string]string {
	filteredParams := make(map[string]string)
	for k, v := range params {
		if v != "" {
			filteredParams[k] = v
		}
	}
	return filteredParams
}

// 企业级签名
func EsMd5(params map[string]string) string {
	serialized := Serialize(params)
	return Sig(serialized)
}

// 安全部签名
func WgMD5(method string, params map[string]string, contentType string) string {
	isPost := strings.ToLower(method) == "post"
	if isPost && contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	filteredParams := FilterUndefinedParams(params)
	// 模拟实现 getSign，实际签名生成需要SDK支持
	return GetSign(filteredParams, contentType, isPost)
}

// 模拟getSign签名生成函数
func GetSign(params map[string]string, contentType string, isPost bool) string {
	// 此处模拟返回MD5签名，可用真实签名逻辑替换
	data := Serialize(params) + contentType
	return Sig(data)
}

// 添加签名到请求参数
func Tranver(params map[string]string) map[string]string {
	params["sig"] = EsMd5(params)
	return params
}

// 获取当前时间戳（毫秒）
func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

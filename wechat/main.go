package main

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
	stdlog "log"
)

func main() {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache

	//init config
	appId := "yours appId"
	AppSecret := "yours AppSecret"
	Token := ""
	EncodingAESKey := ""
	offCfg := &offConfig.Config{
		AppID:          appId,
		AppSecret:      AppSecret,
		Token:          Token,
		EncodingAESKey: EncodingAESKey,
		Cache:          cache.NewMemory(),
		UseStableAK:    false,
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount := wc.GetOfficialAccount(offCfg)

	template := officialAccount.GetTemplate()

	sendData := buildData()

	templateMessage := message.TemplateMessage{
		ToUser:      "o9SiT7FArSDT455bXgiuLYYdHm2A",
		TemplateID:  "EAZetTHKL2YwAeSm2cfp5x7L1Ss3sE0CBr3aKx-xd_A",
		URL:         "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=12312&fenlei=256&rsv_pq=0xc9ca4df00002bb7a&rsv_t=e7ad%2Bf5UlT7m483x3gmB8feQCmW37z5Blx2V9qnJKBZJH2UvjAnwGEqCZfnb&rqlang=en&rsv_dl=tb&rsv_enter=0&rsv_sug3=6&rsv_sug1=1&rsv_sug7=001&rsv_btype=i&prefixsug=12%2526lt%253B12&rsp=5&rsv_sug9=es_0_1&inputT=1037&rsv_sug4=1233&rsv_sug=9",
		Color:       "",
		Data:        sendData,
		ClientMsgID: "",
	}
	send, err := template.Send(&templateMessage)
	stdlog.Println(send, err)
}

func buildData() map[string]*message.TemplateDataItem {
	res := make(map[string]*message.TemplateDataItem)
	one := message.TemplateDataItem{
		Value: "1234567",
		Color: "",
	}
	res["OrderId"] = &one
	res["hotelName"] = &one
	return res
}

//func main() {
//	errMsg := "123"
//	fmt.Println(errors.Is(errors.New(errMsg), errors.New(errMsg))) //结果反直觉，结果为false
//
//	err123 := errors.New(errMsg)
//	fmt.Println(errors.Is(err123, err123)) //结果为true
//}

// package main
//
// import (
//
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//
// )
//
// func main() {
//
//		url := "https://esapp.xiaojukeji.com/hotel-intl/cn/getRoomRatePlanList/v1?wsgsig=dd05-4ltmD5eCeHMw8tPfWUf9DEVonZFvJE5rND%2FR9MeDlVPF8PnF4jqeIgwY4aKdeMBKCgs9rnUzatAiK0HJUE3cCFTwtYUq5aBl%2BGTw9HBEk%2FkI82ia%2BDdgIgYy3kDb9PA3GDmdlbTodGeY3tkNVFWbCl5ql6ESI9eq2WlPaY9HsAa12PXINtV9MW%2Fr1rvabPL%2BDGt2mgLwCj6mKc5cqA8gbh9VyJwT8cFIcsYEEPExYnUS55zGGdEKbXSNBdil%23%230HWLaMhPWIdyvAOBzSV79nl9OReba5WHebntbdwwIj3MfZgSJyBJw88wwLXQP9AZH41yw%2Bc5amm7IqGom4J4Tg%3D%3D"
//		method := "POST"
//		queryStr := "app_id=esapp_pc&datatype=101&encrypt_method=md5&key=8m6cj858wyqjfkfkveab7dtnppjmsf5e&source_channel=200070&app_version=3.3.0&order_source=10&perday_money_quota=9999999900&travel_type=0&city_id=1&hotel_id=3138769&didi_hotel_id=3138769&checkin_date=2024-11-12&checkout_date=2024-11-13&agreement_priority=0&call_type=%5B1%2C3%5D&quick_code=&interest_id=&rp_key=13745123-1-66895644%7C0006-66895644-113-0006-0103-2-0-6401--1&institution_id=1125915971783539&requisition_id=0&hotel_price=31500&is_support_new_multi_standards=1&timestamp=1731381996081&uid=862024277417803&token=DAY3-PXLf0IR746pRhxMI69Mnj-kSOklhLHEWY08SGc0zDkOwjAQhtG7fPUo-sd27DAtPXdgCUtjJBBVlLsjItG-4i10EeRBgzC6E270RHjy3KrRMyGjF8JbnpS9lR-PBPsDxpEA40RMNSmV1FrxNikbF8JlzMTC-_l5nef_uxrXLVQaxy28EUjymopq3WHcCRzjQWj9BgAA__8%3D&program_version=2.13.89&app_time=1731381996&sig=ffac9a51670cd2a843ded816056acc21"
//		payload := strings.NewReader(queryStr)
//
//		params := ParseQueryParams(queryStr)
//		log.Println(params["sig"])
//		delete(params, "sig")
//		filteredParams := WgMD5("post", params, "application/x-www-form-urlencoded")
//
//		log.Println("-------------------")
//		log.Println(filteredParams)
//		log.Println("-------------------")
//
//		client := &http.Client{}
//		req, err := http.NewRequest(method, url, payload)
//
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		req.Header.Add("Accept", "application/json, text/plain, */*")
//		req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
//		req.Header.Add("Cache-Control", "no-cache")
//		req.Header.Add("Connection", "keep-alive")
//		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//		req.Header.Add("Origin", "https://client.es.xiaojukeji.com")
//		req.Header.Add("Pragma", "no-cache")
//		req.Header.Add("Referer", "https://client.es.xiaojukeji.com/")
//		req.Header.Add("Sec-Fetch-Dest", "empty")
//		req.Header.Add("Sec-Fetch-Mode", "cors")
//		req.Header.Add("Sec-Fetch-Site", "same-site")
//		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
//		req.Header.Add("didi-header-hint-content", "{\"es-company-id\":\"1125910870947166\",\"Cityid\":\"-1\"}")
//		req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"130\", \"Google Chrome\";v=\"130\", \"Not?A_Brand\";v=\"99\"")
//		req.Header.Add("sec-ch-ua-mobile", "?0")
//		req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
//
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer res.Body.Close()
//
//		body, err := io.ReadAll(res.Body)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(string(body))
//	}
//
// // ParseQueryParams 将查询字符串转化为 map[string]string
//
//	func ParseQueryParams(query string) map[string]string {
//		values, err := url.ParseQuery(query)
//		if err != nil {
//			fmt.Println("Error parsing query:", err)
//			return nil
//		}
//
//		params := make(map[string]string)
//		for k, v := range values {
//			if len(v) > 0 {
//				params[k] = v[0]
//			}
//		}
//		return params
//	}
package main

import "strings"

func main() {
	split := strings.Split("上海  北京  深圳  广州  杭州  成都  苏州  武汉  西安  重庆  合肥  南京  郑州  长沙  宁波  青岛  无锡  天津  东莞  济南  常州  厦门  南昌  福州  沈阳  珠海  昆明  嘉兴  盐城  惠州  石家庄  金华  烟台  南通  温州  佛山  芜湖  长春  南宁  太原  贵阳  哈尔滨  海口  大连  宁德  保定  潍坊  湖州  徐州  临沂  眉山  泉州  兰州  台州  绍兴  株洲  乌鲁木齐  襄阳  中山  洛阳  扬州  滁州  淮安  西宁  湛江  吉安  宜昌  济宁  赣州  连云港  银川  廊坊  唐山  呼和浩特  宜宾  柳州  江门  鄂尔多斯  泰州  沧州  日照  肇庆  上饶  邯郸  淄博  阜阳  新乡  邢台  绵阳  汕头  包头  河源  漳州  桂林  衢州  宿迁  遵义  乐山  咸阳  十堰  枣庄  安庆  晋中  镇江  抚州  南阳  荆州  汕尾  蚌埠  平顶山  宜春  泰安  宣城  衡阳  榆林  九江  秦皇岛  黔南布依族苗族自治州  商丘  德州  周口  信阳  安阳  聊城  菏泽  揭阳  毕节  常德  晋城  开封  清远  荆门  邵阳  三亚  大理白族自治州  新余  孝感  驻马店  德阳  曲靖  威海  东营  南充  雅安  滨州  铜陵  临汾  泸州  湘潭  茂名  张家口  大同  黄冈  延安  渭南  鄂州  韶关  六安  丽水  许昌  红河哈尼族彝族自治州  宝鸡  伊犁哈萨克自治州  运城  淮南  潮州  长治  宿州  黄石  遂宁  衡水  益阳  岳阳  赤峰  漯河  池州  承德  萍乡  梅州  鞍山  阿拉善盟  安顺  阿里地区  安康  阿克苏地区  阿勒泰地区  阿拉尔  阿德莱德  爱知县  本溪  北海  亳州  巴彦淖尔  白山  白城  百色  白沙黎族自治县  巴中  保山  白银  巴音郭楞蒙古自治州  博尔塔拉蒙古自治州  保亭黎族苗族自治县  布里斯班  北海道  郴州  巢湖市  朝阳  崇左  澄迈县  楚雄彝族自治州  昌都  昌吉回族自治州  昌江黎族自治县  冲绳  大庆  丹东  大兴安岭地区  儋州  东方  定安县  达州  德宏傣族景颇族自治州  迪庆藏族自治州  定西  东京  大阪  大分  恩施土家族苗族自治州  抚顺  阜新  防城港  福冈  贵港  广元  广安  甘孜藏族自治州  甘南藏族自治州  固原  果洛藏族自治州  宫城  冈山  广岛县  葫芦岛  汉中  黄山  淮北  呼伦贝尔  兴安盟  鹤岗  黑河  怀化  鹤壁  贺州  河池  海东  海北藏族自治州  黄南藏族自治州  海南藏族自治州  海西蒙古族藏族自治州  哈密  和田地区  黄金海岸  霍巴特  和歌山  锦州  吉林  景德镇  焦作  鸡西  佳木斯  济源  金昌  嘉峪关  酒泉  吉朗  静冈  京都  喀什地区  克拉玛依  克孜勒苏柯尔克孜自治州  辽阳  莱芜  丽江  拉萨  吕梁  辽源  龙岩  娄底  来宾  临高县  凉山彝族自治州  六盘水  临沧  林芝  陇南  临夏回族自治州  陵水黎族自治县  乐东黎族自治县  马鞍山  牡丹江  墨尔本地区  南平  内江  阿坝藏族羌族自治州  怒江傈僳族自治州  那曲  纽卡斯尔  奈良  攀枝花  莆田  盘锦  濮阳  普洱  平凉  珀斯  齐齐哈尔  七台河  潜江  钦州  琼海  黔西南布依族苗族自治州  黔东南苗族侗族自治州  庆阳  琼中黎族苗族自治县  青森县  秋田  群马  埼玉  千叶  绥化  四平  石河子  朔州  松原  双鸭山  三明  随州  神农架林区  三门峡  山南  日喀则  商洛  石嘴山  三沙  神奈川  兵库县  山口  铁岭  天水  通辽  通化  天门  屯昌县  铜仁  铜川  吐鲁番  塔城地区  中国台湾北部  乌海  乌兰察布  梧州  五指山  文昌  万宁  文山壮族苗族自治州  武威  吴忠  忻州  锡林郭勒盟  咸宁  仙桃  湘西土家族苗族自治州  西双版纳傣族自治州  中国香港  悉尼  新泻  熊本  营口  阳泉  延边朝鲜族自治州  玉林  伊春  鹰潭  永州  阳江  云浮  玉溪  玉树藏族自治州  阳光海岸  舟山  张家界  自贡  资阳  昭通  张掖  中卫  杨凌示范区  中国澳门  图木舒克  五家渠  北屯  铁门关  双河  可克达拉  昆玉  胡杨河  台湾省中部  台湾省南部  台湾省东部  台湾省云林嘉义  台湾省澎湖  中国台湾高雄  中国台湾花莲县  中国台湾嘉义  中国台湾嘉义县  中国台湾基隆  中国台湾金门  中国台湾马祖  中国台湾苗栗县  中国台湾南投县  中国台湾澎湖县  中国台湾屏东县  中国台湾台北  中国台湾台东县  中国台湾台南  中国台湾台中  中国台湾桃园  中国台湾新北  中国台湾新竹  中国台湾新竹县  中国台湾宜兰县  中国台湾云林县  中国台湾彰化县", "  ")
	println(split)
	println(len(split))
}

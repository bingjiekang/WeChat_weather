package weather

import (
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"WeChat_weather/weather_route"
	"encoding/json"
	"net"
	"net/http"

	"github.com/thinkeridea/go-extend/exnet"
)

// 模版{"nums":6,"cityid":"101040100","city":"重庆","date":"2023-08-01","week":"星期二","update_time":"17:07","wea":"雨","wea_img":"yu","tem":"25","tem_day":"34","tem_night":"26","win":"南风","win_speed":"2级","win_meter":"11km\/h","air":"56","pressure":"972","humidity":"96%"}

/*
"city":"重庆",
*/

// 用来获取ip地址和天气信息及返回
func Index(w http.ResponseWriter, req *http.Request) {
	// 获取访问者ip
	ip := GetIp(w, req)
	fmt.Println(ip)
	// fmt.Println(GetWeather(ip))
	// 写到网页界面上
	view := GetWeather(ip)
	w.Write([]byte(view))

}

// 获取访问者的ip地址
func GetIp(w http.ResponseWriter, req *http.Request) string {
	// 获取访问者的ip
	remoteAddr := req.RemoteAddr
	if ip := exnet.ClientPublicIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := exnet.ClientIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// 获取指定ip地址的天气
func GetWeather(ip string) string {
	// 用来获取指定ip天气的url
	url := fmt.Sprintf("https://v0.yiketianqi.com/free/day?appid=81665891&appsecret=kfMA249f&unescape=1&ip=%s", ip)
	// 请求url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取天气失败", err)
	}
	// 对url请求流关闭
	defer resp.Body.Close()
	// 获取信息的全部信息
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败", err)
		return ""
	}
	// 结构体用来接受信息和选择展示指定信息
	var msg weatherroute.Msg
	json.Unmarshal(body, &msg)
	fmt.Println(string(body))
	fmt.Println(msg.City, msg.Wea)

	return string(body)

}

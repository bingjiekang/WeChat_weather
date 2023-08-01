package main

import (
	"WeChat_weather/weather"
	// "github.com/gorilla/mux"

	"net/http"
)

func main() {
	// route := mux.NewRouter()
	// 回调处理
	http.HandleFunc("/", weather.Index)
	http.HandleFunc("/wx", weather.WxVerify)
	// 监听端口
	http.ListenAndServe(":8082", nil)

}

package weatherroute

// "nums":21,"cityid":"101010100","city":"北京","date":"2023-08-01","week":"星期二","update_time":"18:47","wea":"阴","wea_img":"yin","tem":"26","tem_day":"28","tem_night":"25","win":"南风","win_speed":"1级","win_meter":"5km\/h","air":"56","pressure":"1002","humidity":"93%"
type Msg struct {
	Nums        int    `json:nums`        // 今日实时请求次数
	Cityid      string `json:cityid`      // 城市id
	City        string `json:city`        // 城市名称
	Date        string `json:data`        // 当天日期
	Week        string `json:week`        // 当天周几
	Update_time string `json:update_time` // 更新时间
	Wea         string `json:wea`         // 天气状况
	Wea_img     string `json:wea_img`     // 天气标识
	Tem         string `json:tem`         // 实况温度
	Tem_day     string `json:tem_day`     // 白天温度(高温)
	Tem_night   string `json:tem_night`   // 夜间温度(低温)
	Win         string `json:win`         // 风向
	Win_speed   string `json:win_speed`   // 风力
	Win_meter   string `json:win_meter`   // 风速
	Air         string `json:air`         // 空气质量
	Pressure    string `json:pressure`    // 气压
	Humidity    string `json:humidity`    // 湿度
}

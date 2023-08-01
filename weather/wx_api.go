package weather

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
)

func WxVerify(w http.ResponseWriter, req *http.Request) {
	// 1.获取4个字段
	signature := req.URL.Query().Get("signature")
	timestamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	echostr := req.URL.Query().Get("echostr")
	// 2.赋值token
	token := "123456789"
	// 3.token，timestamp，nonce按字典排序的字符串list
	strs := sort.StringSlice{token, timestamp, nonce} // 使用微信公众上的token生成校验
	sort.Strings(strs)
	// 存储排序后的字符串
	ltstr := ""
	for _, v := range strs {
		ltstr += v
	}
	// 4. 哈希算法加密list得到hashcode
	h := sha1.New()
	h.Write([]byte(ltstr))
	hashcode := fmt.Sprintf("%x", h.Sum(nil)) // h.Sum(nil) 做hash
	// 判断hashcode和signature是否相同
	if hashcode == signature { // 相同验证成功,返回echostr
		w.Write([]byte(echostr))
	} else {
		fmt.Println("验证微信公众号失败!")
	}
	// fmt.Println("以下是信息")
	// fmt.Println(signature, timestamp, nonce, echostr, token)
}

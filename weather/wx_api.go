package weather

import (
	"fmt"
	"net/http"
)

func WxVerify(w http.ResponseWriter, req *http.Request) {
	signature := req.Form["signature"]
	timestamp := req.Form["timestamp"]
	nonce := req.Form["nonce"]
	echostr := req.Form["echostr"]
	token := "123456789"
	fmt.Println("以下是信息")
	fmt.Println(signature, timestamp, nonce, echostr, token)
}

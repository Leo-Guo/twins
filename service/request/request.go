package request

import (
	"github.com/Leo-Guo/twins/components/http"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

type RpcConfig struct {
	Uri            string
	Timeout        time.Duration
	WaringDuration time.Duration
}

const DefaultTimeout = 3 * time.Second

func Send(service string, url string, method string, req interface{}, res interface{}) (error, []byte) {
	// 序列化 post_body
	reqBody, errMarshal := jsoniter.Marshal(req)
	if errMarshal != nil {
		return errMarshal, nil
	}
	if strings.EqualFold(method, "POST") {
		response, errReq := http.PostJson(service, url, reqBody, DefaultTimeout)
		if errReq != nil {
			return errReq, nil
		}
		if response.StatusCode == 500 {
			return errReq, response.Body
		}
		if response.StatusCode == 404 {
			return errReq, response.Body
		}
		// 反序列化 response_body
		errUnMarshal := jsoniter.Unmarshal(response.Body, res)
		if errUnMarshal != nil {
			return errReq, response.Body
		}
		return errReq, response.Body
	}
	return nil, []byte{}
}

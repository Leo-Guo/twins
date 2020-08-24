package request

import (
	"github.com/Leo-Guo/twins/common"
	"github.com/Leo-Guo/twins/components/http"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type RpcConfig struct {
	Uri            string
	Timeout        time.Duration
	WaringDuration time.Duration
}

const DefaultTimeout = 3 * time.Second

func Send(service string,url string,method string,req interface{},res interface{}) (common.ErrorCode, []byte){
	// 序列化 post_body
	reqBody, errMarshal := jsoniter.Marshal(req)
	if errMarshal != nil {
		return common.ERROR_MARSHAL, nil
	}
	if method == "POST"{
		response, errReq := http.PostJson(service, url, reqBody, DefaultTimeout)
		if errReq != nil {
			return common.ERROR_REQUEST_UNKNOW, nil
		}
		if response.StatusCode == 500 {
			return common.ERROR_REQUEST_INTERNAL_ERROR_500, response.Body
		}
		if response.StatusCode == 404 {
			return common.ERROR_REQUEST_INTERNAL_ERROR_500, response.Body
		}
		// 反序列化 response_body
		errUnMarshal := jsoniter.Unmarshal(response.Body, res)
		if errUnMarshal != nil {
			return common.ERROR_UNMARSHAL, response.Body
		}
		return common.ERROR_OK, response.Body
	}
	return common.ERROR_REQUEST_UNKNOW, []byte{}
}
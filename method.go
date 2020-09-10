package twins

import (
	"github.com/Leo-Guo/twins/common"
	"github.com/Leo-Guo/twins/service/compare"
	"github.com/Leo-Guo/twins/service/request"
	"github.com/kataras/iris/v12"
)



func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func SendRequest(service string,url string,method string,req interface{},res interface{}) (common.ErrorCode, []byte){
	return  request.Send(service, url, method, req, res)
}

func CompareJson(jsonA map[string]interface{}, jsonB map[string]interface{}, n int)   (string, bool){
	return  compare.JsonCompareDiff(jsonA,jsonB,n)
}

func Out() string{
	return "hello"
}
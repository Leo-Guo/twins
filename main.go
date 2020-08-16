package main

import "github.com/kataras/iris/v12"



func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func out() string{
	return "hello"
}
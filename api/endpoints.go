package api

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)
const (
	BASE_URL = "https://krisha.kz"
)

func GetRent() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		status,body,err := fasthttp.Get([]byte(""),BASE_URL+"/arenda/kvartiry/")
		if err!=nil {
			writeResponse(ctx,status,[]byte(err.Error()))
		}

		res,err  := parseBody(string(body))
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte(err.Error()))
		}
		data,err := json.Marshal(res)
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte(err.Error()))
		}
		writeResponse(ctx,status,data)

	}
}

func GetSell() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		status,body,err := fasthttp.Get([]byte(""),BASE_URL+"/prodazha/kvartiry/")
		if err!=nil {
			writeResponse(ctx,status,[]byte(err.Error()))
		}

		res,err  := parseBody(string(body))
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte(err.Error()))
		}
		data,err := json.Marshal(res)
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte(err.Error()))
		}
		writeResponse(ctx,status,data)
	}
}





func writeResponse(ctx *fasthttp.RequestCtx, status int, msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}
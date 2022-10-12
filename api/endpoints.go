package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	BASE_URL = "https://krisha.kz"
)

func GetRent() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		status, body, err := fasthttp.Get([]byte(""), BASE_URL+"/arenda/kvartiry/")
		if err != nil {
			writeResponse(ctx, status, []byte(err.Error()))
		}

		res, err := parseBody(string(body))
		if err != nil {
			writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		}
		data, err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		}

		writeResponse(ctx, status, data)

	}
}

func GetSell() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		status, body, err := fasthttp.Get([]byte(""), BASE_URL+"/prodazha/kvartiry/")
		if err != nil {
			writeResponse(ctx, status, []byte(err.Error()))
		}

		res, err := parseBody(string(body))
		if err != nil {
			writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		}
		data, err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		}
		writeResponse(ctx, status, data)
	}
}

func GetRentByParams() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		getByParams(ctx, "/arenda/kvartiry")
	}
}

func GetSellByParams() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		getByParams(ctx, "/prodazha/kvartiry")
	}
}

var params = map[string]string{
	"room":       "das[live.rooms]",
	"price_from": "das[price][from]",
	"price_to":   "das[price][to]",
	"has_photo":  "das[_sys.hasphoto]",
	"checked":    "das[checked]",
	"owner":      "das[who]",
	"building":   "das[flat.building]", // 1 кирпичный, 2 панельный, 3 монолитный, 0 иное
	"floor_from": "das[flat.floor][from]",
	"floor_to":   "das[flat.floor][to]",
	"year_from":  "das[house.year][from]",
	"year_to":    "das[house.year][to]",
	"toilet":     "das[flat.toilet]",    // 1 раздельный, 2 совмещенный,3) 2 с/у и более, 4 нет
	"priv_dorm":  "das[flat.priv_dorm]", // 1 yes. 2 no
	"page":       "page",
}

func getByParams(ctx *fasthttp.RequestCtx, link string) {

	city := ""
	form := ""

	pars := make(map[string]string)
	for k, _ := range params {
		b := ctx.QueryArgs().Peek(k)
		if b != nil {
			pars[k] = string(b)
		}
	}

	for k, v := range pars {
		if k == "city" {
			city += v
		}
		form += fmt.Sprintf("%v=%v&", params[k], v)
	}
	requrl := ""
	if city != "" {
		requrl += BASE_URL + link + "/" + city + "/?" + form
	} else {
		requrl += BASE_URL + link + "?" + form
	}
	status, body, err := fasthttp.Get([]byte(""), requrl)
	if err != nil {
		writeResponse(ctx, status, []byte(err.Error()))
		return
	}
	res, err := parseByParams(string(body))

	if err != nil {
		writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		return
	}
	data, err := json.Marshal(res)
	if err != nil {
		writeResponse(ctx, fasthttp.StatusInternalServerError, []byte(err.Error()))
		return
	}
	writeResponse(ctx, fasthttp.StatusOK, data)
}

func writeResponse(ctx *fasthttp.RequestCtx, status int, msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}

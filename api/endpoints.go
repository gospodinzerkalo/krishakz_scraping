package api

import (
	"encoding/json"
	"fmt"
	"strings"

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
		vars := fmt.Sprintf("%v", ctx.UserValue("params"))
		getByParams(ctx, vars, "/arenda/kvartiry")
	}
}

func GetSellByParams() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		vars := fmt.Sprintf("%v", ctx.UserValue("params"))
		getByParams(ctx, vars, "/prodazha/kvartiry")
	}
}

func getByParams(ctx *fasthttp.RequestCtx, par, link string) {
	vars := par

	city := ""
	// replace parameters for get request
	params := make(map[string]string, 0)
	params["room"] = "das[live.rooms]"
	params["price_from"] = "das[price][from]"
	params["price_to"] = "das[price][to]"
	params["has_photo"] = "das[_sys.hasphoto]"
	params["checked"] = "das[checked]"
	params["owner"] = "das[who]"
	params["building"] = "das[flat.building]" // 1 кирпичный, 2 панельный, 3 монолитный, 0 иное
	params["floor_from"] = "das[flat.floor][from]"
	params["floor_to"] = "das[flat.floor][to]"
	params["year_from"] = "das[house.year][from]"
	params["year_to"] = "das[house.year][to]"
	params["toilet"] = "das[flat.toilet]"       // 1 раздельный, 2 совмещенный,3) 2 с/у и более, 4 нет
	params["priv_dorm"] = "das[flat.priv_dorm]" // 1 yes. 2 no
	params["page"] = "page"

	form := ""

	list := strings.Split(vars, "&")

	for _, v := range list {
		spl := strings.Split(v, "=")
		key, val := spl[0], spl[1]
		if key == "city" {
			city += val
		}
		form += fmt.Sprintf("%v=%v&", params[key], val)
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

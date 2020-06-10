package api

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"errors"
)

func parseBody(body string) ([]*Result,error) {

	doc,err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err!=nil {
		return nil,err
	}

	// highlighted section in page
	highlighted := doc.Find(".highlighted-section")

	// first highlighted
	title1,_ := highlighted.Find("div:nth-child(3) .a-card__inc a picture img").Attr("title")
	alt1,_ := highlighted.Find("div:nth-child(3) .a-card__inc a picture img").Attr("alt")
	price1 :=strings.TrimSpace(highlighted.Find("div:nth-child(3) .a-card__descr .a-card__header .a-card__price").Text())
	region1 :=strings.TrimSpace(highlighted.Find("div:nth-child(3) .a-card__descr .a-card__header-body .a-card__wrapper-subtitle div:nth-child(1)").Text())
	preview1 :=strings.TrimSpace(highlighted.Find("div:nth-child(3) .a-card__descr .a-card__header-body .a-card__text-preview").Text())
	link1,_ :=highlighted.Find("div:nth-child(3) .a-card__descr .a-card__header .a-card__main-info .a-card__header-left a").Attr("href")
	first := &Result{
		Title: title1,
		Alt:   alt1,
		Price: price1,
		Region: region1,
		Preview: preview1,
		Link: BASE_URL+link1,
	}

	// second highlighted
	title2,_ := highlighted.Find("div:nth-child(4) .a-card__inc a picture img").Attr("title")
	alt2,_ := highlighted.Find("div:nth-child(4) .a-card__inc a picture img").Attr("alt")
	price2 :=strings.TrimSpace(highlighted.Find("div:nth-child(4) .a-card__descr .a-card__header .a-card__price").Text())
	region2 :=strings.TrimSpace(highlighted.Find("div:nth-child(4) .a-card__descr .a-card__header-body .a-card__wrapper-subtitle div:nth-child(1)").Text())
	preview2 :=strings.TrimSpace(highlighted.Find("div:nth-child(4) .a-card__descr .a-card__header-body .a-card__text-preview").Text())
	link2,_ :=highlighted.Find("div:nth-child(4) .a-card__descr .a-card__header .a-card__main-info .a-card__header-left a").Attr("href")
	second := &Result{
		Title: title2,
		Alt:   alt2,
		Price: price2,
		Region: region2,
		Preview: preview2,
		Link: BASE_URL+link2,
	}

	// third highlighted
	title3,_ := highlighted.Find("div:nth-child(5) .a-card__inc a picture img").Attr("title")
	alt3,_ := highlighted.Find("div:nth-child(5) .a-card__inc a picture img").Attr("alt")
	price3 :=strings.TrimSpace(highlighted.Find("div:nth-child(5) .a-card__descr .a-card__header .a-card__price").Text())
	region3 :=strings.TrimSpace(highlighted.Find("div:nth-child(5) .a-card__descr .a-card__header-body .a-card__wrapper-subtitle div:nth-child(1)").Text())
	preview3 :=strings.TrimSpace(highlighted.Find("div:nth-child(5) .a-card__descr .a-card__header-body .a-card__text-preview").Text())
	link3,_ :=highlighted.Find("div:nth-child(5) .a-card__descr .a-card__header .a-card__main-info .a-card__header-left a").Attr("href")
	third := &Result{
		Title: title3,
		Alt:   alt3,
		Price: price3,
		Region: region3,
		Preview: preview3,
		Link: BASE_URL+link3,
	}


	res := []*Result{first,second,third}

	// all other kv
	all := doc.Find(".a-list.a-search-list.a-list-with-favs")
	allRes,err := parseOther(all)
	if err!=nil {
		return res,err
	}
	res= append(res,allRes...)
	return res,nil
}


func parseByParams(body string) ([]*Result,error){

	not_found := strings.Index(body,"Увы, таких объявлений нет") >= 0
	if not_found {
		return nil,errors.New("Not Found")
	}
	return parseBody(body)

}



// parse all other
func parseOther(doc *goquery.Selection) ([]*Result,error) {
	res := make([]*Result,0)
	doc.Find("div").Each(func(i int, selection *goquery.Selection) {
		title,_ := selection.Find(".a-card__inc a picture img").Attr("title")
		if len(title)==0{
			return
		}
		alt,_ := selection.Find(".a-card__inc a picture img").Attr("alt")
		price :=strings.TrimSpace(selection.Find(".a-card__descr .a-card__header .a-card__price").Text())
		region :=strings.TrimSpace(selection.Find(".a-card__descr .a-card__header-body .a-card__wrapper-subtitle div:nth-child(1)").Text())
		preview :=strings.TrimSpace(selection.Find(".a-card__descr .a-card__header-body .a-card__text-preview").Text())
		link,_ :=selection.Find(".a-card__descr .a-card__header .a-card__main-info .a-card__header-left a").Attr("href")
		item := &Result{
			Title: title,
			Alt:   alt,
			Price: price,
			Region: region,
			Preview: preview,
			Link: BASE_URL+link,
		}
		if len(res)==0 {
			res = append(res,item)
		}
		if res[len(res)-1].Title != title{
			res = append(res,item)
		}
		return
	})
	return res,nil
}

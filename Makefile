build:
	go build main.go
run:
	./main start

depends:
	go get "github.com/valyala/fasthttp"
	go get "github.com/buaazp/fasthttprouter"
	go get "github.com/urfave/cli"
	go get "github.com/PuerkitoBio/goquery"
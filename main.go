package main

import (
	"fmt"
	"log"
	"os"
	"github.com/buaazp/fasthttprouter"
	"github.com/gospodinzerkalo/krishakz_scraping/api"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

func main() {

	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the local server",
			Action: StartServer,
		},
	}
	app.Run(os.Args)

}

func StartServer(d *cli.Context) error {
	router := fasthttprouter.New()

	//endpoints...

	router.GET("/sell", api.GetSell())
	router.GET("/rent", api.GetRent())
	router.GET("/sell/:params", api.GetSellByParams())
	router.GET("/rent/:params", api.GetRentByParams())
	log.Fatal(fasthttp.ListenAndServe(GetPort(), router.Handler))
	return nil
}
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

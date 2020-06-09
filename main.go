package main

import (
	"./api"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"github.com/urfave/cli"
	"log"
	"os"
	"fmt"
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
	router.GET("/allRent",api.GetRent())
	router.GET("/allSell",api.GetSell())
	router.GET("/sell/:params",api.GetSellByParams())
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
package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

var (
	version    = "0.0.1"
	configPath = "./config/"
)

type Data struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Config struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Data    string `json:"data"`
}

type Configs []Config

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "etcc"
	app.Usage = "A simple etc center"
	app.Author = "maemual"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: configPath,
			Usage: "The path of all config",
		},
	}

	app.Action = func(c *cli.Context) {
		if _, err := os.Stat(c.String("path")); err != nil && !os.IsExist(err) {
			os.MkdirAll(c.String("path"), os.ModePerm)
		}
		configPath = c.String("path")

		router := NewRouter()
		log.Fatal(http.ListenAndServe(":8009", router))
	}

	app.Run(os.Args)
}

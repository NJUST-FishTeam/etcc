package main

import (
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	version    = "0.0.1"
	configPath = "./config/"
)

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

		gmux := mux.NewRouter()
		gmux.HandleFunc("/{service}/{config}", getConfigHandler).Methods("GET")
		gmux.HandleFunc("/{service}/{config}", postConfigHandler).Methods("POST")

		http.Handle("/", gmux)

		http.ListenAndServe(":8009", nil)
	}

	app.Run(os.Args)
}

func getConfigHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := params["service"]
	config := params["config"]
	bytes, err := ioutil.ReadFile(filepath.Join(configPath, service, config+".json"))
	if err != nil {
		http.NotFound(w, r)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}

func postConfigHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := params["service"]
	config := params["config"]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server Error", 504)
	}
	err = ioutil.WriteFile(filepath.Join(configPath, service, config+".json"), bytes, os.ModePerm)
	if err != nil {
		http.Error(w, "Server Error", 504)
	}
	w.Write([]byte("OK"))
}

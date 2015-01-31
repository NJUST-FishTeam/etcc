package main

import (
	"encoding/json"
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

type Data struct {
	Data interface{} `json:"data"`
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

		gmux := mux.NewRouter()
		gmux.HandleFunc("/{service}/{config}", getConfigHandler).Methods("GET")
		gmux.HandleFunc("/{service}/{config}", postConfigHandler).Methods("POST")
		gmux.HandleFunc("/", getServiceHandler).Methods("GET")
		gmux.HandleFunc("/{service}", getConfigsHandler).Methods("GET")

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
	err = ioutil.WriteFile(filepath.Join(configPath, service, config+".json"), bytes, 0644)
	if err != nil {
		http.Error(w, "Server Error", 504)
	}
	w.Write([]byte("OK"))
}

func getServiceHandler(w http.ResponseWriter, r *http.Request) {
	configFile, err := os.Open(configPath)
	if err != nil {
		http.Error(w, "Server Error", 504)
	}
	fileInfos, err := configFile.Readdir(1000)
	services := make([]string, 0)
	for _, f := range fileInfos {
		if f.IsDir() {
			services = append(services, f.Name())
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Data{
		Data: services,
	})
}

func getConfigsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	serviceFile := params["service"]
	configFile, err := os.Open(filepath.Join(configPath, serviceFile))
	if err != nil {
		http.Error(w, "Server Error", 504)
	}
	fileInfos, err := configFile.Readdir(1000)
	configs := make([]string, 0)
	for _, f := range fileInfos {
		if !f.IsDir() {
			configs = append(configs, f.Name())
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Data{
		Data: configs,
	})
}

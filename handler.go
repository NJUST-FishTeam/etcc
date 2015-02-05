package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := params["service"]
	config := params["config"]
	bytes, err := ioutil.ReadFile(filepath.Join(configPath, service, config+".json"))
	if err != nil {
		http.NotFound(w, r)
	}
	c := Config{
		Name:    config,
		Service: service,
		Data:    string(bytes),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(c)
}

func PostConfigHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := params["service"]
	config := params["config"]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = ioutil.WriteFile(filepath.Join(configPath, service, config+".json"), bytes, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("OK"))
}

func GetServicesHandler(w http.ResponseWriter, r *http.Request) {
	configFile, err := os.Open(configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func GetConfigsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	serviceFile := params["service"]
	configFile, err := os.Open(filepath.Join(configPath, serviceFile))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fileInfos, err := configFile.Readdir(1000)
	configs := make(Configs, 0)
	for _, f := range fileInfos {
		if !f.IsDir() {
			config := Config{
				Name:    strings.TrimSuffix(f.Name(), ".json"),
				Service: params["service"],
			}
			bytes, _ := ioutil.ReadFile(filepath.Join(configPath, serviceFile, f.Name()))
			config.Data = string(bytes)
			configs = append(configs, config)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(configs)
}

func NewServiceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	service := r.FormValue("service")
	servicePath := filepath.Join(configPath, service)
	if _, err := os.Stat(servicePath); err != nil && !os.IsExist(err) {
		os.MkdirAll(servicePath, os.ModePerm)
	}
	w.Write([]byte("OK"))
}

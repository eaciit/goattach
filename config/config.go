package config

import (
	_ "fmt"
	"strings"
	//"github.com/eaciit/goattach/server/nginx"
	"encoding/json"
	"io/ioutil"
	_ "path/filepath"
)

const serverjson = "E:\\Workspace\\src\\github.com\\eaciit\\goattach\\config\\config.json"

type Config struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Dosync   string `json:"dosync"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Template string `json:"template"`
}

func InitConfig(server string) (Server, error) {
	var rserver Server
	var c Config

	data, err := ioutil.ReadFile(serverjson)
	if err != nil {
		return Server{}, err
	}

	if err = json.Unmarshal(data, &c); err != nil {
		return Server{}, err
	}

	for i := 0; i < cap(c.Servers); i++ {
		if strings.EqualFold(server, c.Servers[i].Name) {
			rserver := c.Servers[i]
			_ = rserver
			return rserver, err
		}
	}

	return rserver, err
}

package config

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Config struct {
	Goattach GoAttach `json:"goattach"`
}

type GoAttach struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Dosync   string `json:"dosync"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Template string `json:"template"`
}

// Uri get the full URI of the configuration file
func (s *Server) Uri() string {
	return filepath.Join(s.Path, s.Filename)
}

func InitConfig(server string, serverjson string) (Server, error) {
	var rserver Server
	var c Config

	data, err := ioutil.ReadFile(serverjson)
	if err != nil {
		return Server{}, err
	}

	if err = json.Unmarshal(data, &c); err != nil {
		return Server{}, err
	}

	for i := 0; i < cap(c.Goattach.Servers); i++ {
		if strings.EqualFold(server, c.Goattach.Servers[i].Name) {
			rserver := c.Goattach.Servers[i]
			_ = rserver
			return rserver, err
		}
	}

	return rserver, err
}

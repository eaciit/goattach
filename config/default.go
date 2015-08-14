package config

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"path/filepath"
)

type Default struct {
	ConfigJson    string `json:"configjson"`
	AppConfigJson string `json:"appconfigjson"`
	Path          string `json:"path"`
}

func InitDefault() (Default, error) {
	defaultjson, err := filepath.Abs("../goattach/config/default.json")

	var dflt Default

	data, err := ioutil.ReadFile(defaultjson)

	if err != nil {
		return Default{}, err
	}

	if err = json.Unmarshal(data, &dflt); err != nil {
		return Default{}, err
	}

	if dflt != (Default{}) {
		dflt.Path, err = filepath.Abs("../goattach/" + dflt.Path)
		dflt.AppConfigJson = filepath.FromSlash(dflt.Path + "/" + dflt.AppConfigJson)
		dflt.ConfigJson = filepath.FromSlash(dflt.Path + "/" + dflt.ConfigJson)
	}

	return dflt, err
}

package config

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// App Object for creating the server paramenter
type App struct {
	Alias string `json:"alias"`
	Appid string `json:"appid"`
	Path  string `json:"path"`
	Port  int    `json:"port"`
	Root  string `json:"root"`
}

type AppList struct {
	Apps []App `json:"app"`
}

// User object for user login information
type User struct {
	UserName string
	Password string
}

// filename, to get the filename of the conf file
func (a *App) filename() string {
	return a.Appid + ".conf"
}

// Uri get the full URI of the configuration file
func (a *App) Uri() string {
	return filepath.Join(a.Path, a.filename())
}

// UpdateList will update the apps.json list
func UpdateList(a App, appsjson string) error {
	list, err := initApps(appsjson)

	if err != nil {
		return err
	}

	exist := false

F:
	for i := 0; i < len(list.Apps); i++ {
		if strings.EqualFold(list.Apps[i].Appid, a.Appid) {
			exist = true
			break F
		}
	}

	if !exist {
		list.Apps = append(list.Apps, a)
		if err := writeAppsJson(list, appsjson); err != nil {
			return err
		}
	}

	return nil
}

// RemoveList will remove the apps from apps.json list
func RemoveList(a App, appsjson string) error {
	list, err := initApps(appsjson)

	if err != nil {
		return err
	}

F:
	for i := 0; i < len(list.Apps); i++ {
		if strings.EqualFold(list.Apps[i].Appid, a.Appid) {
			list.Apps = append(list.Apps[:i], list.Apps[i+1:]...)
			if err := writeAppsJson(list, appsjson); err != nil {
				return err
			}
			break F
		}
	}

	return nil
}

func initApps(appsjson string) (AppList, error) {

	var list AppList

	data, err := ioutil.ReadFile(appsjson)
	if err != nil {
		return AppList{}, err
	}

	if err = json.Unmarshal(data, &list); err != nil {
		return AppList{}, err
	}

	return list, err
}

func writeAppsJson(list AppList, appsjson string) error {
	result, err := json.Marshal(list)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(appsjson, result, 0644); err != nil {
		return err
	}

	return nil
}

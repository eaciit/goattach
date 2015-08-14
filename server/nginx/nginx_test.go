package nginx

import (
/*"encoding/json"
"fmt"
_ "github.com/eaciit/config"
//"os"
"io/ioutil"
"testing"*/
)

/*const path = "E:\\Workspace\\src\\github.com\\eaciit\\goattach\\config\\config.json"

type Config struct {
	Servers []Server `json:"servers"`
	Default Default  `json:"default"`
}

type Server struct {
	Dosync   string `json:"dosync"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Template string `json:"template"`
}

type Default struct {
	Defaultserver string `json:"defaultserver"`
}

func TestGetJson(t *testing.T) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var c Config

	if err = json.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	fmt.Println(c.Default.Defaultpath)
	fmt.Println(c.Default.Defaultserver)

	fmt.Println(c.Servers[0].Filename)
	fmt.Println(c.Servers[0].Dosync)
	fmt.Println(c.Servers[0].Name)
	fmt.Println(c.Servers[0].Path)
	fmt.Println(c.Servers[0].Template)

	fmt.Println(c.Servers[1].Filename)
	fmt.Println(c.Servers[1].Dosync)
	fmt.Println(c.Servers[1].Name)
	fmt.Println(c.Servers[1].Path)
	fmt.Println(c.Servers[1].Template)

	fmt.Println("OK")
}*/

/*const path = "E:\\Workspace\\src\\github.com\\eaciit\\goattach\\config\\apps.json"

type App struct {
	Alias string `json:"alias"`
	Appid string `json:"appid"`
	Path  string `json:"path"`
	Port  string `json:"port"`
	Root  string `json:"root"`
}

type AppList struct {
	Apps []App `json:"app"`
}

func TestGetJson(t *testing.T) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var r AppList

	if err = json.Unmarshal(data, &r); err != nil {
		panic(err)
	}

	fmt.Println(r.Apps[0].Alias)
	fmt.Println(r.Apps[0].Appid)
	fmt.Println(r.Apps[0].Path)
	fmt.Println(r.Apps[0].Port)
	fmt.Println(r.Apps[0].Root)

	fmt.Println(r.Apps[1].Alias)
	fmt.Println(r.Apps[1].Appid)
	fmt.Println(r.Apps[1].Path)
	fmt.Println(r.Apps[1].Port)
	fmt.Println(r.Apps[1].Root)

	fmt.Println("OK")
}*/

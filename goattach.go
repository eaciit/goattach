package main

import (
	"flag"
	"fmt"
	cfg "github.com/eaciit/goattach/config"
	nginx "github.com/eaciit/goattach/server/nginx"
	"path/filepath"
	"strings"
)

const (
	NOK           = "NOK"
	OK            = "OK"
	DefaultServer = "nginx"
)

var rserver cfg.Server
var rdefault cfg.Default

func init() {
	var err error
	rdefault, err = cfg.InitDefault()
	if err != nil {
		fmt.Println(NOK + " " + err.Error())
	}
}

func main() {
	attach := flag.Bool("attach", false, "command for attach the new configuration file")
	appId := flag.String("id", NOK, "Id for the configuration file")
	alias := flag.String("alias", NOK, "Alias for the app")
	port := flag.Int("port", 0, "Port for the app")
	sync := flag.Bool("sync", false, "Synchronize the server")
	detach := flag.Bool("detach", false, "Detach the config file")
	server := flag.String("webserver", DefaultServer, "Type of the server")
	path := flag.String("path", rdefault.Path, "Path to put the configuration file")
	root := flag.String("root", "", "Root config")
	config := flag.String("config", "", "Server config")
	appconfig := flag.String("appconfig", "", "App config")

	flag.Parse()

	app := cfg.App{
		Alias: *alias,
		Appid: *appId,
		Path:  filepath.Clean(*path),
		Port:  *port,
		Root:  *root,
	}

	if *config != "" {
		rdefault.ConfigJson = filepath.Clean(*config)
	}

	if *appconfig != "" {
		rdefault.AppConfigJson = filepath.Clean(*appconfig)
	}

	if !strings.EqualFold(*server, DefaultServer) {
		rserver = initConfig(*server, rdefault.ConfigJson)
	} else {
		rserver = initConfig(DefaultServer, rdefault.ConfigJson)
	}

	success := true

	if *attach && success {

		if app.Appid == NOK || app.Alias == NOK || app.Port == 0 {
			fmt.Println(NOK + " Please provide the complete data for attach")
		} else {
			switch *server {
			case "nginx":

				if err := nginx.WriteConf(app, rserver); err != nil {
					fmt.Println(NOK + " " + err.Error())
					success = false
				}
				if err := cfg.UpdateList(app, rdefault.AppConfigJson); err != nil {
					fmt.Println(NOK + " " + err.Error())
					success = false
				}

				if success {
					fmt.Println(OK + " Attach into nginx ")
				}
			case "apache":
				fmt.Println(OK + " Attach into apache ")
			default:
				success = false
				fmt.Println(NOK + " Unrecognized server type ")
			}
		}

	}

	if *detach && success {
		if app.Appid == NOK {
			fmt.Println(NOK + " Please provide the complete data for detach")
		} else {
			switch *server {
			case "nginx":

				if err := nginx.RemoveConf(app, rserver); err != nil {
					success = false
					fmt.Println(NOK + " " + err.Error())
				}
				if err := cfg.RemoveList(app, rdefault.AppConfigJson); err != nil {
					success = false
					fmt.Println(NOK + " " + err.Error())
				}

				if success {
					fmt.Println(OK + " Detach from nginx")
				}
			case "apache":

				fmt.Println(OK + " Detach from apache")
			default:

				fmt.Println(NOK + " Unrecognized server type ")
			}
		}
	}

	if *sync && success {
		switch *server {
		case "nginx":
			if err := nginx.Sync(rserver.Dosync); err != nil {
				fmt.Println(NOK + " Synchronize " + err.Error())
			} else {
				fmt.Println(OK + " Synchronize nginx server")
			}
		case "apache":
			fmt.Println(OK + " Synchronize apache server")
		default:
			fmt.Println(NOK + " Unrecognized server type ")
		}
	}
}

func initConfig(server string, configjson string) cfg.Server {
	rserver, err := cfg.InitConfig(server, configjson)
	if err != nil {
		fmt.Println(NOK + " " + err.Error())
	}

	return rserver
}

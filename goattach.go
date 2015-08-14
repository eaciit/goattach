package main

import (
	"flag"
	"fmt"
	config "github.com/eaciit/goattach/config"
	nginx "github.com/eaciit/goattach/server/nginx"
)

const (
	NOK = "NOK"
	OK  = "OK"
)

func main() {
	attach := flag.Bool("attach", false, "command for attach the new configuration file")
	appId := flag.String("id", NOK, "Id for the configuration file")
	alias := flag.String("alias", NOK, "Alias for the app")
	port := flag.Int("port", 0, "Port for the app")
	sync := flag.Bool("sync", false, "Synchronize the server")
	detach := flag.Bool("detach", false, "Detach the config file")
	server := flag.String("webserver", "nginx", "Type of the server")
	path := flag.String("path", NOK, "Path to put the configuration file")
	root := flag.String("root", "", "Root config")

	flag.Parse()

	app := config.App{
		Alias: *alias,
		Appid: *appId,
		Path:  *path,
		Port:  *port,
		Root:  *root,
	}

	rserver, err := config.InitConfig(*server)

	if err != nil {
		panic(err)
	}

	if *attach {
		if app.Appid == NOK || app.Alias == NOK || app.Port == 0 || app.Path == NOK {
			fmt.Println(NOK + "Please provide the complete data")
		} else {
			switch *server {
			case "nginx":
				fmt.Println("Attach into nginx ")
				if err := nginx.WriteConf(app, rserver); err != nil {
					panic(err)
				}
				if err := config.UpdateList(app); err != nil {
					panic(err)
				}
			case "apache":
				fmt.Println("Attach into apache")
			default:
				fmt.Println(NOK + " unrecognized server type ")
			}
		}

	}

	if *detach {
		if app.Appid == NOK || app.Path == NOK {
			fmt.Println(NOK)
		} else {
			switch *server {
			case "nginx":
				fmt.Println("Detach from nginx")

				if err := nginx.RemoveConf(app, rserver); err != nil {
					panic(err)
				}
				if err := config.RemoveList(app); err != nil {
					panic(err)
				}

			case "apache":
				fmt.Println("Detach from apache")
			default:
				fmt.Println(NOK + " unrecognized server type ")
			}
		}
	}

	if *sync {
		switch *server {
		case "nginx":
			fmt.Println("Synchronize nginx server")
			if err := nginx.Sync(rserver.Dosync); err != nil {
				//panic(err)
			}
		case "apache":
			fmt.Println("Synchronize apache server")
		default:
			fmt.Println(NOK + " unrecognized server type ")
		}
	}
}

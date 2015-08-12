package main

import (
	"fmt"
	. "github.com/eaciit/goattach/conf"
)

func main() {
	w := WebApp{
		AppId:                "test",
		Path:                 "E:\\Workspace\\src\\github.com\\eaciit\\goattach\\conffile",
		Listen:               "7070",
		ServerName:           "www.goattach.com",
		DefaultConfDirectory: "conffile",
		ProxyPass:            "http://127.0.0.1",
		// other parameters
	}

	if e := WriteConf(w); e != "" {
		panic(e)
	} else {
		fmt.Println("Success!")
	}

	//fmt.Println(RemoveConf(w))
}

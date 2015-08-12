package conf

import (
	_ "fmt"
	"path/filepath"
)

type WebApp struct {
	AppId                string
	Path                 string
	Listen               string
	ServerName           string
	ProxyPass            string
	DefaultConfDirectory string
}

type User struct {
	UserName string
	Password string
}

func (w *WebApp) filename() string {
	return w.AppId + ".conf"
}

func (w *WebApp) Uri() string {
	return filepath.Join(w.Path, w.filename())
}

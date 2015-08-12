package conf

import (
	_ "fmt"
	"path/filepath"
)

// WebApp Object for creating the server paramenter
type WebApp struct {
	AppId                string
	Path                 string
	Listen               string
	ServerName           string
	ProxyPass            string
	DefaultConfDirectory string
}

// User object for user login information
type User struct {
	UserName string
	Password string
}

// filename, to get the filename of the conf file
func (w *WebApp) filename() string {
	return w.AppId + ".conf"
}

// Uri get the full URI of the configuration file
func (w *WebApp) Uri() string {
	return filepath.Join(w.Path, w.filename())
}

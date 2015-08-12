package conf

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	//"reflect"
)

const defaultConfFileName = "nginx.conf"

var defaultConfFile string
var fileUri string
var filePth string

// createDir, create the directory for the conf file
func createDir() error {
	if err := os.MkdirAll(filePth, 0777); err != nil {
		return err
	}
	return nil
}

// write, write the content to the file
func write(s []byte, confFile string) error {
	if err := ioutil.WriteFile(confFile, s, 0644); err != nil {
		return err
	}
	return nil
}

// template will form the server conf content
func template(w WebApp) string {
	strTemplate := "server {"

	if w.Listen != "" {
		strTemplate += "\n\tlisten " + w.Listen + ";"
	}

	if w.ServerName != "" {
		strTemplate += "\n\tserver_name " + w.ServerName + ";"
	}

	if w.ProxyPass != "" {
		strTemplate += "\n\tlocation /{\n\t\tproxy_pass\t" + w.ProxyPass + "\n\t}"
	}

	strTemplate += "\n}"

	return strTemplate
}

// WriteConf is the main function for creating the conf file
func WriteConf(w WebApp) string {
	fileUri = w.Uri()
	filePth = w.Path

	if err := isDefaultConfExist(w); err != true {
		return "Default Conf File Error: file is not exist"
	}

	result := []byte(template(w))

	if err := createDir(); err != nil {
		return "Create Directory Error: " + err.Error()
	}

	if err := write(result, fileUri); err != nil {
		return "Conf File Write Error: " + err.Error()
	}

	if isConfigFileAdded(w) == false {
		if err := attachConfFile(w); err != nil {
			return "Attach Conf File Error: " + err.Error()
		}

		reloadServer()
	}

	return ""
}

// RemoveConf is to remove the conf file
func RemoveConf(w WebApp) string {
	if err := os.Remove(fileUri); err != nil {
		return "Error Remove the File: " + err.Error()
	}

	return ""
}

// isDefaultConfExist will check the default conf file
func isDefaultConfExist(w WebApp) bool {
	_, err := os.Stat(filepath.Join(w.DefaultConfDirectory, defaultConfFileName))
	defaultConfFile = filepath.Join(w.DefaultConfDirectory, defaultConfFileName)
	return os.IsNotExist(err) == false
}

// isConfigFileAdded will check the default conf file
// if the conf file already include in the default conf file, then no need to include the conf file
func isConfigFileAdded(w WebApp) bool {
	result := string(readDefaultConfFile(w))
	return strings.Contains(result, fileUri)
}

// attachConfFile function is to attach the conf file into the default conf file
func attachConfFile(w WebApp) error {

	result := string(readDefaultConfFile(w))
	lines := strings.Split(result, "\n")
	linesLength := cap(lines)

F:
	for i := linesLength - 1; i > 0; i-- {
		if strings.TrimSpace(lines[i]) == "}" {
			lines[i] = getIncludeTemplate() + "\n}"
			break F
		}
	}

	output := strings.Join(lines, "\n")

	err := ioutil.WriteFile(defaultConfFile, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

// readDefaultConfFile function is to read the default conf file
func readDefaultConfFile(w WebApp) []byte {
	result, err := ioutil.ReadFile(filepath.Join(w.DefaultConfDirectory, defaultConfFileName))

	if err != nil {
		panic(err)
	}

	return result
}

// getIncludeTemplate function is to get the include formatting with the conf file name
func getIncludeTemplate() string {
	return "\n\tinclude\t" + fileUri
}

// reloadServer function is to reload the nginx server
// to apply the new settings, the server need to reload first
func reloadServer() error {
	cmd := exec.Command("nginx -s reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

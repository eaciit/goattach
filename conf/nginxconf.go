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

func load() error {
	var err error = nil
	if isConfigFileExist(fileUri) == false {
		if err := createDir(); err != nil {
			return err
		}
		if err := ioutil.WriteFile(fileUri, []byte(""), 0644); err != nil {
			return err
		}
	}

	return err
}

func createDir() error {
	if err := os.MkdirAll(filePth, 0777); err != nil {
		return err
	}
	return nil
}

func write(s []byte, confFile string) error {
	if err := ioutil.WriteFile(confFile, s, 0644); err != nil {
		return err
	}
	return nil
}

func isConfigFileExist(confFile string) bool {
	_, err := os.Stat(confFile)
	return os.IsNotExist(err) == false
}

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

func WriteConf(w WebApp) string {
	fileUri = w.Uri()
	filePth = w.Path

	if err := load(); err != nil {
		return "Conf File Error: " + err.Error()
	}

	if err := isDefaultConfExist(w); err != true {
		return "Default Conf File Error: file is not exist"
	}

	//fmt.Println(string(readDefaultConfFile(w)))

	result := []byte(template(w))

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

func RemoveConf(w WebApp) string {
	if err := os.Remove(fileUri); err != nil {
		return "Error Remove the File: " + err.Error()
	}

	return ""
}

func isDefaultConfExist(w WebApp) bool {
	_, err := os.Stat(filepath.Join(w.DefaultConfDirectory, defaultConfFileName))
	defaultConfFile = filepath.Join(w.DefaultConfDirectory, defaultConfFileName)
	return os.IsNotExist(err) == false
}

func isConfigFileAdded(w WebApp) bool {
	result := string(readDefaultConfFile(w))
	return strings.Contains(result, fileUri)
}

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

func readDefaultConfFile(w WebApp) []byte {
	result, err := ioutil.ReadFile(filepath.Join(w.DefaultConfDirectory, defaultConfFileName))

	if err != nil {
		panic(err)
	}

	return result
}

func getIncludeTemplate() string {
	return "\n\tinclude\t" + fileUri
}

func reloadServer() error {
	cmd := exec.Command("nginx -s reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

package nginx

import (
	"errors"
	_ "fmt"
	config "github.com/eaciit/goattach/config"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	//"reflect"
)

const NOK = "NOK"
const OK = "OK"

// createDir, create the directory for the conf file
func createDir(a config.App) error {
	if err := os.MkdirAll(a.Path, 0777); err != nil {
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

func createTemplate(s config.Server, a config.App) string {
	var result string
	r, err := ioutil.ReadFile(filepath.Join("./template", s.Template))

	if err != nil {
		panic(err)
	}

	result = string(r)

	result = strings.Replace(result, "@port", strconv.Itoa(a.Port), 1)
	result = strings.Replace(result, "@alias", a.Alias, 1)
	result = strings.Replace(result, "@root", a.Root, 1)

	return result
}

// WriteConf is the main function for creating the conf file
func WriteConf(a config.App, s config.Server) error {
	if err := isDefaultConfExist(s); err != true {
		return errors.New(NOK + " Default Conf File Error: file is not exist")
	}
	result := []byte(createTemplate(s, a))

	if err := createDir(a); err != nil {
		return err //return NOK + "Create Directory Error: " + err.Error()
	}

	if err := write(result, a.Uri()); err != nil {
		return err //return NOK + "Conf File Write Error: " + err.Error()
	}

	if err := attachConfFile(a, s); err != nil {
		return err //return NOK + "Conf File Attach Error: " + err.Error()
	}

	return nil
}

// RemoveConf is to remove the conf file
func RemoveConf(a config.App, s config.Server) error {
	if err := os.Remove(a.Uri()); err != nil {
		return err //return NOK + "Error Remove the File: " + err.Error()
	}

	detachConfFile(a, s)

	return nil
}

// isDefaultConfExist will check the default conf file
func isDefaultConfExist(s config.Server) bool {
	_, err := os.Stat(filepath.Join(s.Path, s.Filename))
	return os.IsNotExist(err) == false
}

// isConfigFileAdded will check the default conf file
// if the conf file already include in the default conf file, then no need to include the conf file
func isConfigFileAdded(a config.App, s config.Server) bool {
	result := string(readDefaultConfFile(s))
	return strings.Contains(result, a.Uri())
}

// attachConfFile function is to attach the conf file into the default conf file
func attachConfFile(a config.App, s config.Server) error {

	result := string(readDefaultConfFile(s))
	lines := strings.Split(result, "\n")
	linesLength := len(lines)

	added := strings.Contains(result, a.Uri())

	if !added {

	F:
		for i := linesLength - 1; i > 0; i-- {
			if strings.TrimSpace(lines[i]) == "}" {
				lines[i] = getIncludeTemplate(a) + "\n}"
				break F
			}
		}

		output := strings.Join(lines, "\n")

		err := ioutil.WriteFile(filepath.Join(s.Path, s.Filename), []byte(output), 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

func detachConfFile(a config.App, s config.Server) error {

	result := string(readDefaultConfFile(s))
	lines := strings.Split(result, "\n")
	linesLength := len(lines)

	added := strings.Contains(result, a.Uri())

	if added {

	F:
		for i := linesLength - 1; i > 0; i-- {
			if strings.TrimSpace(lines[i]) == strings.TrimSpace(getIncludeTemplate(a)) {
				lines[i] = ""
				break F
			}
		}

		output := strings.Join(lines, "\n")

		err := ioutil.WriteFile(filepath.Join(s.Path, s.Filename), []byte(output), 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

// readDefaultConfFile function is to read the default conf file
func readDefaultConfFile(s config.Server) []byte {
	result, err := ioutil.ReadFile(filepath.Join(s.Path, s.Filename))

	if err != nil {
		panic(err)
	}

	return result
}

// getIncludeTemplate function is to get the include formatting with the conf file name
func getIncludeTemplate(a config.App) string {
	return "\n\tinclude\t" + a.Uri()
}

// reloadServer function is to reload the nginx server
// to apply the new settings, the server need to reload first
func Sync(c string) error {
	cmd := exec.Command(c)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

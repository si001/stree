package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"si001/stree/model"
)

var fileErr error
var settingsDir string
var settingsFileName string

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadSettings() {
	settingsDir = getHomeDir() + model.PathDivider + "." + model.AppName
	if !exists(settingsDir) {
		fileErr = os.Mkdir(settingsDir, os.ModePerm)
	}
	settingsFileName = settingsDir + model.PathDivider + "settings"
	if exists(settingsFileName) {
		jsData, err := ioutil.ReadFile(settingsFileName)
		if err == nil {
			err = json.Unmarshal(jsData, &model.DataSettings)
			if err != nil {
				fmt.Print("error read settings", err)
			}
		}
	}
}

func WriteSettings() {
	jsData, _ := json.Marshal(model.DataSettings)
	err := ioutil.WriteFile(settingsFileName, jsData, 0644)
	if err != nil {
		fmt.Print("error write settings", err)
	}
}

func ReadHistory(t string) []string {
	fn := settingsDir + model.PathDivider + t + ".hst"
	if exists(fn) {
		hd, err := ioutil.ReadFile(fn)
		if err == nil {
			var rd []string
			err = json.Unmarshal(hd, &rd)
			if err != nil {
				fmt.Print("error read "+fn, err)
			}
			return rd
		}
	}
	return []string{""}
}

func WriteHistory(t string, hist []string) {
	fn := settingsDir + model.PathDivider + t + ".hst"
	hd, _ := json.Marshal(hist)
	err := ioutil.WriteFile(fn, hd, 0644)
	if err != nil {
		fmt.Print("error write "+fn, err)
	}
}

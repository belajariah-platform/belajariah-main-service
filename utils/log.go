package utils

import (
	"belajariah-main-service/config"
	"belajariah-main-service/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func InitializeCustomLogs(config *model.LogConfig) error {
	isExist := fileExists(config.Filename)
	if !isExist {
		_, err := os.Create(config.Filename)
		return err
	}
	return nil
}

func PushLog(v ...interface{}) {
	fileName := fmt.Sprintf("./%s", config.GetConfig().Log.Filename)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}
	today := time.Now().In(loc)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}

	loging := model.Logs{}
	err = json.Unmarshal(data, &loging)

	loging.Logs = append(loging.Logs, model.LogsData{
		Date: today.Format("2006-01-02 15:04:05"),
		Log:  fmt.Sprint(v...),
	})

	bytes, _ := json.Marshal(loging)

	ioutil.WriteFile(fileName, bytes, 0644)
}

func PushLogf(log string, v ...interface{}) {
	fileName := fmt.Sprintf("./%s", config.GetConfig().Log.Filename)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}
	today := time.Now().In(loc)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}

	loging := model.Logs{}
	err = json.Unmarshal(data, &loging)

	jsons, _ := json.Marshal(v)
	str := ""
	if string(jsons) == "null" {
		str = strings.ReplaceAll(fmt.Sprintf("%s", log), "\"", " ")
	} else {
		str = strings.ReplaceAll(fmt.Sprintf("%s, %s", log, string(jsons)), "\"", " ")
	}

	loging.Logs = append(loging.Logs, model.LogsData{
		Date: today.Format("2006-01-02 15:04:05"),
		Log:  strings.ReplaceAll(str, "\\", ""),
	})

	bytes, _ := json.Marshal(loging)

	ioutil.WriteFile(fileName, bytes, 0644)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PushLogStackTrace(timezone string, traces []string) {
	fileName := fmt.Sprintf("./%s", config.GetConfig().Log.Filename)
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}
	today := time.Now().In(loc)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("[ERROR] => ", err)
	}

	loging := model.Logs{}
	err = json.Unmarshal(data, &loging)

	loging.Logs = append(loging.Logs, model.LogsData{
		Date: today.Format("2006-01-02 15:04:05"),
		Log:  strings.Join(traces, " => "),
	})

	bytes, _ := json.Marshal(loging)

	ioutil.WriteFile(fileName, bytes, 0644)
}

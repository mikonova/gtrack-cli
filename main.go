package main

import (
	_ "flag"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		log.Println("err: couldn't access user config dir")
	}
	configPath := userConfigPath + "gtrack-cli/config"
	logPath := userConfigPath + "gtrack-cli/log"
	trackerPath := userConfigPath + "gtrack-cli/tracker"
	err = os.MkdirAll(userConfigPath, 0644)
	if err != nil {
		log.Println("err: couldn't access user config dir")
	}
	err = os.MkdirAll(logPath, 0644)
	if err != nil {
		log.Println("err: couldn't access user logging dir")
	}
	file, err := os.OpenFile(logPath+"/gtrack.log", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("err: couldn't create logfile")
	}
	log.SetOutput(file)
}

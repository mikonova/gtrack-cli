package main

import (
	"flag"
	_ "flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		log.Println("err: couldn't access user config dir")
	}

	uniformPath := filepath.Join(userConfigPath, "gtrack-cli")

	configPath := filepath.Join(uniformPath, "config")
	logPath := filepath.Join(uniformPath, "log")
	trackPath := filepath.Join(uniformPath, "tracker")

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

	flag.Usage = func() {
		fmt.Println(`
		Usage: gtrack-cli [OPTIONAL ARGS] [ARGUMENT] [ARG FLAG] [VALUES...] [OPTIONAL ARGS]
		
		Arguments:

		help: displays this message
		remote-add <address>: add a new remote to sync list
		remote-del <address>: remove a remote from the sync list
		add -all: add a new file to all the remotes
		add -once <remote> <file|dir>: add a file for a certain remote
		del -all <file|dir>: remove a file or dir from all the remotes
		del -once <remote> <file|dir>: remove a file or directory from one remote
		sync -once <remote>: sync all files for a certain remotes
		sync -all: sync all files for all remotes
		fetch -once <remote>: fetch files or dirs from a specific remote
		fetch -all: fetch files from all remotes`)

		flag.PrintDefaults()
	}
}

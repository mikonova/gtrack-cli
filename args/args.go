package args

import (
	"flag"
	"fmt"
	"gtrack-cli/tracker"
	"net"
	"os"
	"path/filepath"
)

var (
	errExists   = "err: this remote already exists"
	errNoSetup  = "err: Couldn't set up a configuration file for the current remote"
	errNoRemote = "err: incorrect remote name or it doesnt exist"
)

func ParseArgs(configPath, trackPath string) (err error) {
	flagsAllowed := true
	skipRounds := 0
	for k, arg := range os.Args {

		if skipRounds > 0 {
			skipRounds--
			continue
		}

		switch arg {
		case "help":
			flag.Usage()
			return
		case "remote-add":
			var (
				conn        net.Conn
				precisePath = filepath.Join(trackPath, os.Args[k+1])
				file        *os.File
			)

			_, err = os.Stat(precisePath)
			if err == nil {
				fmt.Println(errExists)
				err = nil
				return
			}

			conn, err = net.Dial("tcp", os.Args[k+1])
			if err != nil {
				fmt.Println("err: invalid server address or not connected to a network")
				return
			}
			conn.Close()

			err = os.MkdirAll(precisePath, 0644)
			if err != nil {
				fmt.Println(errNoSetup)
				return
			}
			trackerfile := filepath.Join(precisePath, "trackerfile.txt")
			file, err = os.Create(trackerfile)
			if err != nil {
				fmt.Println(errNoSetup)
				return
			}
			file.Close()
			fmt.Println("info: remote added", conn.LocalAddr().String(), "successfully")
			skipRounds = 1

		case "remote-del":
			precisePath := filepath.Join(trackPath, os.Args[k+1])

			_, err = os.Stat(precisePath)
			if err != nil {
				fmt.Println()
				err = nil
				return
			}

			err = os.RemoveAll(precisePath)
			if err != nil {
				fmt.Println("err: deletion error, check log for detail")
				return
			}
			skipRounds = 1

		case "add":
			switch os.Args[k+1] {
			case "--all", "-all":
				// TODO: implement file/dir adding for all remotes
			case "-once", "--once":
				if (os.Args[k+2] == "" || os.Args[k+3] == "") || (os.Args[k+2] == "" && os.Args[k+3] == "") {
					fmt.Println("err: empty args. Format: add <-once|--once> <remote> <file|directory> ")
					err = nil
					return
				}
				precisePath := filepath.Join(trackPath, os.Args[2])
				_, err = os.Stat(precisePath)
				if err != nil {
					fmt.Println(errNoRemote)
					return
				}

				err = tracker.AddTrackerEntry(precisePath, os.Args[3])
				if err != nil {
					return
				}
				skipRounds = 3
				continue
			default:
				fmt.Println("err: incorrect syntax. Expected: \"add <-once |--once> <remote> <file|directory>\" or \"add <-all|--all> <file|directory>\"")
				err = nil
				return
			}

		case "del":
			if (os.Args[k+1] == "" || os.Args[k+2] == "") || (os.Args[k+1] == "" && os.Args[k+2] == "") {
				fmt.Println("err: empty args. Format: del <remote> <file|directory> ")
				err = nil
				return
			}
			precisePath := filepath.Join(trackPath, os.Args[k+1])
			_, err = os.Stat(precisePath)
			if err != nil {
				fmt.Println(errNoRemote)
				return
			}

			err = tracker.RemoveTrackerEntry(os.Args[k+2], precisePath)
			if err != nil {
				return
			}
			skipRounds = 2
		case "sync":
			switch os.Args[k+1] {
			case "-once", "--once":
				precisePath := filepath.Join(configPath, os.Args[k+2])
				if _, err = os.Stat(precisePath); err != nil {
					fmt.Println(errNoRemote, "Format: \"sync -once <remote>\"")
					return
				}
				// TODO: add syncing logic for one remote
				skipRounds = 2

			case "-all", "--all":
				// TODO: add syncing logic for all the files for all remotes
				skipRounds = 1
			default:
				fmt.Println("err: incorrect format. Allowed args after \"sync\" are \"-all\" or \"-once <remote>\"")
				err = nil
				return
			}
		case "fetch":
			switch os.Args[k+1] {
			case "-once", "--once":
				precisePath := filepath.Join(configPath, os.Args[k+2])
				if _, err = os.Stat(precisePath); err != nil {
					fmt.Println(errNoRemote, "Format: \"fetch -once <remote>\"")
					return
				}
				// TODO: add fetching logic for one remote
				skipRounds = 2
			case "-all", "--all":
				// TODO: add fetching logic for all the files for all remotes
				skipRounds = 1
			default:
				fmt.Println("err: incorrect format. Allowed args after \"fetch\" are \"-all\" or \"-once <remote>\"")
				err = nil
				return
			}

		default:
			if os.Args[k] == "--" {
				flagsAllowed = false
				continue
			}
			if os.Args[k][0:0] == "-" && flagsAllowed {
				flag.CommandLine.Parse(os.Args[k:k])
			}

			fmt.Println("err: no such argument exists")
		}
	}
	return nil
}

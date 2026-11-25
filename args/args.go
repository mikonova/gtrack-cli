package args

import (
	"fmt"
	"gtrack-cli/tracker"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

func ParseArg(configPath, trackPath string) (err error) {
	if os.Args[1][1] == '-' {
		return
	}
	switch os.Args[1] {
	case "help":
		var execPath, execPathPrecise string
		execPath, err = os.Executable()
		if err != nil {
			return
		}
		execPathPrecise, err = filepath.EvalSymlinks(execPath)
		if err != nil {
			return
		}
		cmd := exec.Command(execPathPrecise, "-help")
		err = cmd.Run()
		if err != nil {
			return
		}
	case "remote-add":
		var (
			conn        net.Conn
			precisePath = filepath.Join(trackPath, os.Args[2])
			file        *os.File
		)

		_, err = os.Stat(precisePath)
		if err == nil {
			fmt.Println("err: this remote already exists")
			err = nil
			return
		}

		conn, err = net.Dial("tcp", os.Args[2])
		if err != nil {
			fmt.Println("err: invalid server address")
			return
		}
		conn.Close()

		err = os.MkdirAll(precisePath, 0644)
		if err != nil {
			fmt.Println("err: Couldn't set up a configuration dir for the current remote")
			return
		}
		trackerfile := filepath.Join(precisePath, "trackerfile.txt")
		file, err = os.Create(trackerfile)
		if err != nil {
			fmt.Println("err: Couldn't set up a configuration file for the current remote")
			return
		}
		file.Close()
		fmt.Println("info: remote added", conn.LocalAddr().String(), "successfully")

	case "remote-del":
		precisePath := filepath.Join(trackPath, os.Args[2])

		_, err = os.Stat(precisePath)
		if err != nil {
			fmt.Println("err: this remote doesn't exist")
			err = nil
			return
		}

		err = os.RemoveAll(precisePath)
		if err != nil {
			fmt.Println("err: deletion error, check log for detail")
			return
		}

	case "add":
		if (os.Args[2] == "" || os.Args[3] == "") || (os.Args[2] == "" && os.Args[3] == "") {
			fmt.Println("err: empty args. Format: add <remote> <file | folder> ")
			err = nil
			return
		}
		precisePath := filepath.Join(trackPath, os.Args[2])
		_, err = os.Stat(precisePath)
		if err != nil {
			fmt.Println("err: no such remote found")
			return
		}

		err = tracker.AppendTrackerFile(precisePath, os.Args[3])
		if err != nil {
			return
		}
	case "del":
		if (os.Args[2] == "" || os.Args[3] == "") || (os.Args[2] == "" && os.Args[3] == "") {
			fmt.Println("err: empty args. Format: add <remote> <file | folder> ")
			err = nil
			return
		}
		precisePath := filepath.Join(trackPath, os.Args[2])
		_, err = os.Stat(precisePath)
		if err != nil {
			fmt.Println("err: no such remote found")
			return
		}

		err = tracker.RemoveTrackerEntry(os.Args[3], precisePath)
		if err != nil {
			return
		}

	default:
		if os.Args[2][0:0] == "-" {
			return nil
		}

		fmt.Println("err: no such argument exists")
	}
	return nil
}

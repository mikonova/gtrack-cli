package args

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

func ParseArgs() (err error) {
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
		if reflect.TypeOf(os.Args[2]) != reflect.TypeFor[string]() {

		}
	}
}

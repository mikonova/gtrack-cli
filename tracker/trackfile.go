package tracker

import (
	"fmt"
	"gtrack-cli/uplink/protocol/headers/category"
	"gtrack-cli/uplink/protocol/headers/types"
	"log"
	"net"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

var remoteMap map[string][]string

func AddRemote(remote, trackerPath string) {
	if !isConnValid(remote) {
		errStr := "err: connection string is not valid"
		fmt.Println(errStr)
		log.Fatalln(errStr)
	}
	_, err := toml.DecodeFile(trackerPath, &remoteMap)
	if err != nil {
		fmt.Println("err: couldn't read the tracking file")
		log.Fatalln(err.Error())
	}

	if _, exists := remoteMap[remote]; exists {
		fmt.Println("err: remote already added")
		return
	} else {
		remoteMap[remote] = nil
		file, err := os.Create(trackerPath)
		if err != nil {
			fmt.Println("err: Couldn't open the tracking file")
			log.Fatalln("err:", err.Error())
		}
		defer file.Close()

		encoder := toml.NewEncoder(file)
		err = encoder.Encode(remoteMap)
		if err != nil {
			fmt.Println("err: Couldn't encode data to toml and write")
			log.Fatalln("err:", err.Error())
		}
	}
}

func RemoveRemote(remote, trackerPath string) {
	_, err := toml.DecodeFile(trackerPath, &remoteMap)
	if err != nil {
		fmt.Println("err: couldn't read the tracking file")
		log.Fatalln(err.Error())
	}
	delete(remoteMap, remote)
	file, err := os.Create(trackerPath)
	if err != nil {
		fmt.Println("err: Couldn't open the tracking file")
		log.Fatalln("err:", err.Error())
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(remoteMap)
	if err != nil {
		fmt.Println("err: Couldn't encode data to toml and write")
		log.Fatalln("err:", err.Error())
	}

}

func AddOnce(trackerPath, contentPath string) (err error) {

}

func RemoveOnce(trackerPath, contentPath string) (err error) {

}

func GetFor(trackerPath, remoteName string) []string {

}

func AddToAll(trackerPath, contentPath string) {

}

func RemoveFromAll(contentPath string) {

}

func GetAll(contentPath string) {

}

func isConnValid(connString string) bool {
	if conn, err := net.Dial("tcp", connString); err != nil {
		return false
	} else {
		go func() {
			buf := [2]byte{category.INFO, types.ABORT_CONN}
			conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
			conn.Write(buf[:])
			conn.Close()
		}()
		return true
	}
}

package tracker

import (
	"fmt"
	"log"
	"os"
)

func AppendTrackerFile(fdir, content string) (err error) {
	file, err := os.OpenFile(fdir+".trackerfile", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	_, err = os.Stat(content)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("err:", err.Error())
			log.Fatalln("err:", err.Error())
		}
	}
	_, err = file.WriteString()
	if err != nil {
		return
	}
}

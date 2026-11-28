package cliprints

import (
	"bufio"
	"fmt"
	"os"
)

func GetTrackerFile(trackPath string) {
	fmt.Println("trackerfile is located in ", trackPath)
}

func GetTrackedFiles(trackPath string) (err error) {
	file, err := os.Open(trackPath)
	if err != nil {
		fmt.Println("err: could not open trackerfile. See the log for details")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("List of all entries in the trackerfile")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}

package tracker

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AppendTrackerFile(preciseTrackPath, contentPath string) (err error) {
	_, err = os.Stat(contentPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("err: file", filepath.Base(contentPath), "is not found")
			return
		}
		fmt.Println("err: an unexpected error has occured. See the log for details")
		return
	}
	trackPrecise := filepath.Join(preciseTrackPath, "trackerfile.txt")

	file, err := os.OpenFile(trackPrecise, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err: could not open trackerfile. See the log for details")
		return
	}
	defer file.Close()
	_, err = file.WriteString(contentPath + "\n")
	if err != nil {
		fmt.Println("err: could not write to trackerfile. See the log for details")
		return
	}
	return
}

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

func RemoveTrackerEntry(entry, trackPath string) (err error) {
	file, err := os.Open(trackPath)
	if err != nil {
		fmt.Println("err: could not open trackerfile. See the log for details")
		return
	}
	scanner := bufio.NewScanner(file)
	builder := strings.Builder{}
	fmt.Println("List of all entries in the trackerfile")
	for scanner.Scan() {
		text := scanner.Text()
		if n := strings.Compare(entry, text); n == 0 {
			continue
		}
		builder.WriteString(text)
		builder.WriteRune('\n')
	}
	if err = scanner.Err(); err != nil {
		file.Close()
		return
	}
	file.Close()
	file, err = os.OpenFile(trackPath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("err: could not open trackerfile. See the log for details")
		return
	}
	defer file.Close()
	_, err = file.WriteString(builder.String())
	if err != nil {
		fmt.Println("err: could not remove entry: unable to write to trackfile. See the log for details")
		return
	}
	return
}

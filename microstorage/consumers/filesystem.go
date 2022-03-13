package consumers

import (
	"fmt"
	"log"
	"microstorage"
	"os"
	"path/filepath"
	"time"
)

func ListenFilesystem(delay int64) {
	dirname := microstorage.Config.Credentials.Filesystem
	td := int64(time.Millisecond) * delay
	microstorage.LogMessage("Filesystem consumer started")
	for ;; {
		scanDir(dirname)
		time.Sleep(time.Duration(td))
	}
}


func scanDir(dirname string) (bool) {
	files, err := filepath.Glob(dirname + "/data/*.json")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) > 0 {
		for _, fname := range files {
			data := LoadInputJson(fname)
			data.RawPath, _ = filepath.Abs(dirname + "/files/" + data.Id)
			microstorage.LogMessage(fmt.Sprintf("Filesystem start saving %s", data.Id))
			PutFileIntoStorage(data)
			err := os.Remove(fname)
			if err != nil {
				microstorage.ThrowException(fmt.Sprintln(err))
			}
			microstorage.LogMessage(fmt.Sprintf("Filesystem ends saving %s", data.Id))
		}
		return true
	}
	return false
}

func ListenFilesystemInit() {
	microstorage.CreateDir(microstorage.Config.Credentials.Filesystem + "/data")
	microstorage.CreateDir(microstorage.Config.Credentials.Filesystem + "/files")
	microstorage.CreateDir(microstorage.Config.Credentials.Filesystem + "/output")
}
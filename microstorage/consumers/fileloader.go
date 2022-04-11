package consumers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microstorage"
	"os"
	"path/filepath"
)

var BaseStoragePath = ""

func PutFileIntoStorage(data InputData) {
	//todo set filesize
	//microstorage.ThrowException(fmt.Sprintln(err))
	data_path,file_path := GetNestedPath(data.Id)
	microstorage.CreateDir(filepath.Dir(data_path))
	microstorage.CreateDir(filepath.Dir(file_path))
	var err error
	if len(data.RawFile) == 0 {
		err = os.Rename(data.RawPath, file_path)
	} else {
		err = os.WriteFile(file_path, data.RawFile, 0777)
	}
	if err!=nil {
		microstorage.ThrowException(fmt.Sprintln(err))
	}
	data.RawFile = []byte{}
	file, _ := json.MarshalIndent(data, "", " ")
	err = ioutil.WriteFile(data_path, file, 0777)
	if err!=nil {
		microstorage.ThrowException(fmt.Sprintln(err))
	}
}

func GetNestedPath(id string) (data_path string, file_path string) {
	file_path = BaseStoragePath + "/files/"
	data_path = BaseStoragePath + "/data/"
	runes := []rune(id)
	safeSubstring := string(runes[0:2])
	file_path = file_path + safeSubstring + "/" + id
	data_path = data_path + safeSubstring + "/" + id
	return data_path, file_path
}


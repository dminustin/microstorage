package consumers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microstorage"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type InputData struct {
	Id string
	Mime string
	Ext string
	Transform struct {
		Resize []string
		Archive []string
		Convert []string
	}
	RawPath string
	RawFile []byte
	FileSize int
}

func LoadInputJson(filename string) (result InputData){
	pb := path.Base(filename)
	result.Id = strings.TrimSuffix(pb, filepath.Ext(pb))
	jsonFile, err := os.Open(filename)
	if err != nil {
		microstorage.ThrowException(fmt.Sprintln(err))
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	return result
}
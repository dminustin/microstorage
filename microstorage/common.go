package microstorage

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

func StrToInt64(s string) (int64){
	i, err := strconv.ParseInt(s ,10,64)
	if err!=nil {
		ThrowException("Error Conversion")
	}
	return i
}


func CreateDir(dirname string) {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		result := os.MkdirAll(dirname, 0777)
		if result != nil {
			ThrowException(fmt.Sprintf("Cannot create dir %s (%s)", dirname, result))
		}
	}
}

func IsUUID(s string) (bool) {
	//d3d29d70-1d25-11e3-8591-034165a3a613
	_,e := uuid.Parse(s)
	if e!=nil {
		ThrowException(fmt.Sprintln(e))
	}
	return true
}
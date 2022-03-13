package microstorage

import "os"

func ThrowException(message string) {
	LogMessage(message)
	os.Exit(1)
}

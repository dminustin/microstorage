package microstorage

import "fmt"
import "time"

func LogMessage(message string) {
	fmt.Println(fmt.Sprintf("[%s] %s", time.Now().Format("01-02-2006 15:04:05"), message))
}

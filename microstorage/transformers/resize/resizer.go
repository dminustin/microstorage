package resize

import (
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"microstorage"
	"os"
	"path/filepath"
	"strconv"
)

//got from https://www.golangprograms.com/how-to-create-thumbnail-of-an-image-in-golang.html
func ImgResize(filePath string, w,h int64) ([]byte, error) {
	newName := fmt.Sprintf ("%s/resize/%sx%s/%s",
		filepath.Dir(filePath),
		strconv.Itoa(int(w)),
		strconv.Itoa(int(h)),
		filepath.Base(filePath))
	if _, err := os.Stat(newName); errors.Is(err, os.ErrNotExist) {
		microstorage.CreateDir(filepath.Dir(newName))
		img, err := imaging.Open(filePath)
		if err != nil {
			panic(err)
		}
		imaging.JPEGQuality(100)
		thumb := imaging.Thumbnail(img, int(w), int(h), imaging.CatmullRom)
		fmt.Println(newName)

		err = imaging.Save(thumb, newName+".jpg")
		if err != nil {
			microstorage.ThrowException(fmt.Sprintln(err))
		}
		//Because!
		os.Rename(newName+".jpg", newName)
	}
	return os.ReadFile(newName)
}

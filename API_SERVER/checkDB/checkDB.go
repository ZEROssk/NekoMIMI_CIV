package checkDB

import(
	"os"
	"log"
	"time"
	"strings"
	"io/ioutil"
	"image"

	"main/useDB"
	"main/saveIMG"
)

func readDir(p string) []os.FileInfo {
	files, _ := ioutil.ReadDir(p)
	return files
}

func newIMG(files []os.FileInfo, path string) {
	for _, f := range files {
		FB, _ := os.Open(path+"/"+f.Name())
		defer FB.Close()

		decImg, format, err := image.Decode(FB)
		if err != nil {
			log.Println(err)
			return
		} else {
			ID := strings.Split(f.Name(), "-")[2]
			saveIMG.SaveThumbnail(decImg, f.Name(), format)
			useDB.DBaddImg(ID, f.Name())
		}
	}
}

func CheckDB(path string) {
	const tformat = "2006-01-02 15:04:05"

	r := useDB.DBcheckData()
	if r == 0 {
		newIMG(readDir(path), path)
	} else if r != len(readDir(path)) {
		StStamp := useDB.DBcheckCreatedAt()
		rTStamp, _ := time.Parse(tformat, StStamp)

		files := readDir(path)
		for _, f := range files {
			fTStamp, _ := time.Parse(tformat, f.ModTime().Format(tformat))
			if fTStamp.After(rTStamp) == true {
				newIMG([]os.FileInfo{f}, path)
			}
		}
	}
}


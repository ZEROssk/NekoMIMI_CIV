package checkDB

import(
	"os"
	"io/ioutil"
	"log"
	"time"
	"strings"
	"bytes"
	"net/http"
	"image"

	"main/useDB"
	"main/saveIMG"
)

func readDir(p string) []os.FileInfo {
	files, _ := ioutil.ReadDir(p)
	return files
}

func newIMG(files []os.FileInfo, path string) {
	for i, f := range files {
		fpath := path+f.Name()

		os.Chmod(fpath, 0644)
		FB, _ := os.Open(fpath)
		defer FB.Close()

		bufData, err := ioutil.ReadAll(FB)
		if err != nil {
			log.Println(err)
			return
		}

		buf := bytes.NewBuffer(bufData)
		mimeType := http.DetectContentType(buf.Bytes())
		if mimeType != "image/jpeg" && mimeType != "image/png" {
			continue
		}

		decImg, format, err := image.Decode(buf)
		if err != nil {
			log.Println(err)
			return
		} else {
			ID := strings.Split(f.Name(), "-")[2]
			saveIMG.SaveThumbnail(decImg, f.Name(), format)
			useDB.DBaddImg(ID, f.Name())
		}

		log.Println("Progress: ", i+1, "/", len(files))
	}
}

func CheckDB(path string) {
	const tformat = "2006-01-02 15:04:05"

	r := useDB.DBcheckData()
	if r == 0 {
		log.Printf("Add New Image")
		newIMG(readDir(path), path)
	} else if r != len(readDir(path)) {
		log.Printf("Update New Image")
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


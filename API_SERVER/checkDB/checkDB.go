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

func ReadDir(p string) []os.FileInfo {
	files, _ := ioutil.ReadDir(p)
	return files
}

func NewIMG(files []os.FileInfo, path string) {
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

		log.Println("Add Progress: ", i+1, "/", len(files))
	}
}

func CheckDB(path string) {
	r := useDB.DBcheckData()
	if r == 0 {
		log.Printf("Add New Image")
		NewIMG(ReadDir(path), path)
	} else if r != len(ReadDir(path)) {
		const tFormat = "2001-01-01 11:11:11"
		log.Printf("Update New Image")
		StStamp := useDB.DBcheckCreatedAt()
		rTStamp, _ := time.Parse(tFormat, StStamp)

		files := ReadDir(path)
		for i, f := range files {
			fTStamp, _ := time.Parse(tFormat, f.ModTime().Format(tFormat))
			if fTStamp.After(rTStamp) == true {
				NewIMG([]os.FileInfo{f}, path)
			}

			log.Println("Update Progress: ", i+1, "/", len(files))
		}
	}
}

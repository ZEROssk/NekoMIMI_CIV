package checkDB

import(
	"os"
	"time"
	"strings"
	"io/ioutil"

	"../useDB"
)

func readDir(p string) []os.FileInfo {
	files, _ := ioutil.ReadDir(p)
	return files
}

func insertDB(files []os.FileInfo) {
	for _, f := range files {
		ID := strings.Split(f.Name(), "-")[2]
		useDB.DBaddImg(ID, f.Name())
	}
}

func CheckDB(path string) {
	const tformat = "2006-01-02 15:04:05"

	r := useDB.DBcheckData()
	if r == 0 {
		insertDB(readDir(path))
	} else if r != len(readDir(path)) {
		StStamp := useDB.DBcheckCreatedAt()
		rTStamp, _ := time.Parse(tformat, StStamp)

		files := readDir(path)
		for _, f := range files {
			fTStamp, _ := time.Parse(tformat, f.ModTime().Format(tformat))
			if fTStamp.After(rTStamp) == true {
				insertDB([]os.FileInfo{f})
			}
		}
	}
}


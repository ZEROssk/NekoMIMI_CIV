package checkDB

import(
	"os"
	"regexp"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbNAME	string
	dbUSER	string
	dbPORT	string
	dbTABLE	string
	db		*sql.DB
)

type Data struct {
	Rec	int
}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}

func LoginDB() {
	dbNAME	= getENV("MYSQL_DB")
	dbUSER	= getENV("MYSQL_USER")
	dbPORT	= getENV("MYSQL_PORT")
	dbTABLE	= getENV("MYSQL_TABLE")

	var err error
	db, err = sql.Open("mysql", dbUSER+"@tcp(db:"+dbPORT+")/"+dbNAME)
	if err != nil {
		panic(err.Error())
	}
}

func makeFileList() []string {
	fList := []string{}

	path := "/go/Content/Twitter"
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		fList = append(fList, f.Name())
	}
	return fList
}

func addFirstData() {
	fileList := makeFileList()
	rep := regexp.MustCompile(`\s*-\s*`)

	for _, FName := range fileList {
		ID := rep.Split(FName, -1)
		_, err := db.Exec("INSERT INTO twimg_data (TwiID, FileName) VALUES ('?','?');", ID[2], FName)
		if err != nil {
			panic(err.Error())
		}
	}
}

func CheckDB() {
	LoginDB()
	var v Data
	var r int

	rows := db.QueryRow("SELECT COUNT(*) FROM twimg_data;")
	err := rows.Scan(&v.Rec)
	if err != nil {
		panic(err.Error())
	}
	r = v.Rec
	if r == 0 {
		addFirstData()
	}
}


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

func addFirstData() {
	path := "/go/Content/Twitter"
	rep := regexp.MustCompile(`\s*-\s*`)

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		ID := rep.Split(f.Name(), -1)
		_, err := db.Exec("INSERT INTO twimg_data (TwiID, FileName) VALUES (?,?)", ID[2], f.Name())
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


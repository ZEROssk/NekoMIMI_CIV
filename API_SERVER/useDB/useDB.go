package useDB

import(
	"os"
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
	ID			int
	TwiID		string
	Img			string
	CreatedAt	string
}

type Check struct {
	Rec	int
	StStamp string
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

func DBhome(p string, begin string, end string) ([][]string, int) {
	var v Data
	var m int
	s := [][]string{}

	rows, err := db.Query("SELECT*FROM "+dbTABLE+" LIMIT ?, ?", begin, end)
	if err != nil {
		panic(err.Error())
	} else {
		row := db.QueryRow("SELECT COUNT(*) FROM "+dbTABLE+";")

		err := row.Scan(&v.ID)
		if err != nil {
			panic(err.Error())
		}
		m = v.ID
	}

	for rows.Next() {
		err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		s = append(s, []string{v.TwiID, v.Img})
	}
	return s, m
}

func DBsearch(t string, begin string, end string) ([][]string, int) {
	var v Data
	var m int
	s := [][]string{}

	rows, err := db.Query("SELECT*FROM "+dbTABLE+" WHERE TwiID=? LIMIT ?, ?", t, begin, end)
	if err != nil {
		panic(err.Error())
	} else {
		rows := db.QueryRow("SELECT COUNT(*) FROM "+dbTABLE+" WHERE TwiID=? ", t)

		err := rows.Scan(&v.ID)
		if err != nil {
			panic(err.Error())
		}
		m = v.ID
	}

	for rows.Next() {
		err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		s = append(s, []string{v.TwiID, v.Img})
	}
	return s, m
}

func DBorigin(t string, f string) []string {
	var v Data
	row := db.QueryRow("SELECT*FROM "+dbTABLE+" WHERE TwiID=? AND FileName=?", t, f)

	err := row.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
	if err != nil {
		s := []string{}
		return s
	}
	s := []string{v.TwiID, v.Img}

	return s
}

func DBaddImg(t string, f string) {
	_, err := db.Exec("INSERT INTO "+dbTABLE+" (TwiID, FileName) VALUES (?,?)", t, f)
	if err != nil {
		panic(err.Error())
	}
}

func DBcheckData() int {
	var v Check

	row := db.QueryRow("SELECT COUNT(*) FROM "+dbTABLE+";")
	err := row.Scan(&v.Rec)
	if err != nil {
		panic(err.Error())
	}
	return v.Rec
}

func DBcheckCreatedAt() string {
	var v Check

	row := db.QueryRow("select CreatedAt from "+dbTABLE+" where CreatedAt=(select max(CreatedAt) from "+dbTABLE+")")
	err := row.Scan(&v.StStamp)
	if err != nil {
		panic(err.Error())
	}
	return v.StStamp
}

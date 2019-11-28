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

func DB_home(p string, begin string, end string) ([][]string, int) {
	var v Data
	var m int
	s := [][]string{}

	rows, err := db.Query("SELECT*FROM "+dbTABLE+" LIMIT ?, ?", begin, end)
	if err != nil {
		panic(err.Error())
	} else {
		rows := db.QueryRow("SELECT COUNT(*) FROM "+dbTABLE+";")

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

func DB_search(t string, begin string, end string) ([][]string, int) {
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

func DB_origin(t string, f string) []string {
	var v Data
	rows := db.QueryRow("SELECT*FROM "+dbTABLE+" WHERE TwiID=? AND FileName=?", t, f)

	err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
	if err != nil {
		panic(err.Error())
	}
	s := []string{v.TwiID, v.Img}

	return s
}


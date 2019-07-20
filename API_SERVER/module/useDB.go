package useDB

import(
	."fmt"
	"log"
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

func Login_DB() {
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

func DB_home(p string, begin string, end string) []string {
	rows, err := db.Query("SELECT*FROM twi_data LIMIT ?, ?", begin, end)
	if err != nil {
		panic(err.Error())
	}

	s := []string{}

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s %s\n", v.ID, v.TwiID, v.Img, v.CreatedAt)
		s = append(s, Sprintf("%s %s", v.TwiID, v.Img))
	}
	return s
}

func DB_search(t string, begin string, end string) string {
	rows, err := db.Query("SELECT*FROM twi_data WHERE TwiID=? LIMIT ?, ?", t, begin, end)
	if err != nil {
		panic(err.Error())
	}

	s := ""

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s %s\n", v.ID, v.TwiID, v.Img, v.CreatedAt)
		s += Sprintf("%s %s", v.TwiID, v.Img)
	}
	return s
}

func DB_origin(t string, f string) string {
	var v Data
	rows := db.QueryRow("SELECT*FROM twi_data WHERE TwiID=? AND FileName=?", t, f)

	e := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
	if e != nil {
		panic(e.Error())
	}
	s := Sprintf("%s %s", v.TwiID, v.Img)

	return s
}


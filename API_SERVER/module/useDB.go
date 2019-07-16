package useDB

import(
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
)

type Data struct {
	ID		int
	TwiID	string
	IMG		string
	DATE	string
}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}

func DB_login() *sql.DB {
	dbNAME	= getENV("MYSQL_DB")
	dbUSER	= getENV("MYSQL_USER")
	dbPORT	= getENV("MYSQL_PORT")
	dbTABLE	= getENV("MYSQL_TABLE")

	dblogin := dbUSER+"@tcp(db:"+dbPORT+")/"+dbNAME
	db, err := sql.Open("mysql", dblogin)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func DB_home(p string, begin string, end string) {
	db := DB_login()
	defer db.Close()

	// tes := "SELECT MAX(ID) FROM "+dbTABLE+" LIMIT "+begin", "+end
	tes := "SELECT*FROM "+dbTABLE+" LIMIT "+begin", "+end
	rows, err := db,Query(tes)
	if err != nil {
		panic(err.Error())
	}

	// dbTable := "SELECT*FROM "+dbTABLE
	// rows, err := db.Query(dbTable)
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	log.Println(p)

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.TwiID, &v.IMG, &v.DATE)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s %s\n", v.ID, v.TwiID, v.IMG, v.DATE)
	}

}

func DB_search(twiID string, p string) {
	db := DB_login()
	defer db.Close()

	log.Printf("DB_search\n")
	log.Println(twiID, p)
}

func DB_origin(twiID string, imgID string) {
	db := DB_login()
	defer db.Close()

	log.Printf("DB_origin\n")
	log.Println(twiID, imgID)
}


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
	ID			int
	TwiID		string
	Img			string
	CreatedAt	string
}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}

func Login_DB() *sql.DB {
	dbNAME	= getENV("MYSQL_DB")
	dbUSER	= getENV("MYSQL_USER")
	dbPORT	= getENV("MYSQL_PORT")
	dbTABLE	= getENV("MYSQL_TABLE")

	//connectDB := dbUSER+"@tcp(db:"+dbPORT+")/"+dbNAME
	db, err := sql.Open("mysql", dbUSER+"@tcp(db:"+dbPORT+")/"+dbNAME)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func DB_home(p string, begin string, end string) {
	db := Login_DB()
	defer db.Close()

	rows, err := db.Query("SELECT*FROM twi_data LIMIT ?, ?", begin, end)
	if err != nil {
		panic(err.Error())
	}

	log.Println(p)

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s %s\n", v.ID, v.TwiID, v.Img, v.CreatedAt)
	}

}

func DB_search(t string, p string) {
	db := Login_DB()
	defer db.Close()

	log.Printf("DB_search\n")
	log.Println(t, p)
}

func DB_origin(t string, f string) {
	db := Login_DB()
	defer db.Close()

	rows, err := db.Query("SELECT TWI_ID, FILE_NAME FROM twi_data WHERE TWI_ID='?' AND FILE_NAME='?'", t, f)
	if err != nil {
		panic(err.Error())
	}

	log.Println(t, f)

	var v Data
	err := rows.Scan(&v.ID, &v.TwiID, &v.Img, &v.CreatedAt)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("%d %s %s %s\n", v.ID, v.TwiID, v.Img, v.CreatedAt)
}


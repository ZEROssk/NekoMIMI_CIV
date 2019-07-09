package moduleDB

import(
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	ID		int
	UID		string
	IMG		string
	DATE	string
}

func DB() {
	dbNAME := getENV("MYSQL_DB")
	dbUSER := getENV("MYSQL_USER")
	dbPORT := getENV("MYSQL_PORT")
	dbTABLE := getENV("MYSQL_TABLE")

	dblogin := dbUSER+ "@tcp(db:" +dbPORT+ ")/" +dbNAME

	db, err := sql.Open("mysql", dblogin)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbTable := "SELECT*FROM "+dbTABLE

	rows, err := db.Query(dbTable)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.UID, &v.IMG, &v.DATE)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s %s\n", v.ID, v.UID, v.IMG, v.DATE)
	}

}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}


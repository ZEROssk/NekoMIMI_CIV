package moduleDB

import(
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	ID		int
	IMG		string
	DATE	string
}

func DB() {
	dbNAME := getENV("MYSQL_DB")
	dbUSER := getENV("MYSQL_USER")
	dbPORT := getENV("MYSQL_PORT")

	log.Printf("%s %s %s", dbNAME, dbUSER, dbPORT)

	dblogin := dbUSER+ "@tcp(db:" +dbPORT+ ")/" +dbNAME

	db, err := sql.Open("mysql", dblogin)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT*FROM imgs")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var v Data
		err := rows.Scan(&v.ID, &v.IMG, &v.DATE)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("%d %s %s\n", v.ID, v.IMG, v.DATE)
	}

}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}


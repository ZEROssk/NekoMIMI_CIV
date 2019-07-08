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

	db, err := sql.Open("mysql", "root:root@tcp(db:5300)/imgs_data_db")
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


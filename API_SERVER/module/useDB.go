package moduleDB

import(
	"log"
	"os"

	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DB() {
	dbNAME := getENV("MYSQL_DB")
	dbUSER := getENV("MYSQL_USER")
	dbPORT := getENV("MYSQL_PORT")

	log.Printf("%s %s %s", dbNAME, dbUSER, dbPORT)
}

func getENV(p string) string {
	env := os.Getenv(p)
	return env
}


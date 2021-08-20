package db
import (
    "log"
    "fmt"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var dbs map[string]*sql.DB

func init() {
    dbs =  make(map[string]*sql.DB)
}

func Db(conn string) *sql.DB {

    if dbs[conn] != nil {
        return dbs[conn]
    }
    var connString = os.Getenv(fmt.Sprintf("DB_%s", conn))

    log.Println("connecting " + connString)

    db, err := sql.Open("mysql", connString)

    if err != nil {
        panic(err)
    }
    dbs[conn] = db
    return db
}


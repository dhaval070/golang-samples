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
    log.Println("init db")
    dbs =  make(map[string]*sql.DB)
}

func Db(conn string) *sql.DB {

    if dbs[conn] != nil {
        return dbs[conn]
    }
    var connString = os.Getenv(fmt.Sprintf("DB_%s", conn))

    log.Println("connecting " + connString)

    db, err := sql.Open("mysql", connString)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxIdleTime(5000000000)

    if err != nil {
        panic(err)
    }
    dbs[conn] = db
    return db
}


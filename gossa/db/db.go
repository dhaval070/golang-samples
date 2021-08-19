package db
import (
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Db() *sql.DB {
    if db != nil {
        return db
    }
    var connString = os.Getenv("DB_CONN")

    var err error

    db, err = sql.Open("mysql", connString)
    if err != nil {
        panic(err)
    }
    return db
}


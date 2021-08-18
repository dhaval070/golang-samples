package db
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const connString = "dhaval@tcp(127.0.0.1:3306)/ops"

var db *sql.DB

func Db() *sql.DB {
    if db != nil {
        return db
    }
    var err error

    db, err = sql.Open("mysql", connString)
    if err != nil {
        panic(err)
    }
    return db
}


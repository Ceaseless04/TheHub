package sqldb

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"os"
)

func auth() {
    db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("Success!")
}

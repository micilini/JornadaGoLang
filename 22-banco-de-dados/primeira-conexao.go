package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/jornadagolang")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }

    log.Println("Conexão OK!")
}

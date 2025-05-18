
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

func exec(db *sql.DB, sql string) sql.Result{
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

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

    // Fazendo um SELECT na tabela usuarios
	rows, err := db.Query("SELECT id, nome, email, criado_em FROM usuarios")
	if err != nil {
		log.Fatal("Erro ao buscar usuários:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nome, email string
		var criadoEm string
		err := rows.Scan(&id, &nome, &email, &criadoEm)
		if err != nil {
			log.Println("Erro ao ler linha:", err)
		}
		fmt.Printf("ID: %d | Nome: %s | Email: %s | Criado em: %s\n", id, nome, email, criadoEm)
	}
}

package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

func ExecQuery(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, fmt.Errorf("erro ao executar query: %w", err)
    }
    return result, nil
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

    // Fazendo um UPDATE na tabela usuário
	nomeNovo := "Roll Atualizado"
	id := 1

	_, err = ExecQuery(db, "UPDATE usuarios SET nome = ? WHERE id = ?", nomeNovo, id)
	if err != nil {
		log.Fatal("Erro ao atualizar nome:", err)
	}
	fmt.Println("Nome atualizado com sucesso!")
}
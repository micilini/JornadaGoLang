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

	// Criando sua primeira tabela de usuarios
	query := `
		CREATE TABLE IF NOT EXISTS usuarios (
			id INT AUTO_INCREMENT PRIMARY KEY,
			nome VARCHAR(100),
			email VARCHAR(100),
			criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	_, err = ExecQuery(db, query)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	log.Println("Tabela 'usuarios' criada com sucesso!")

}
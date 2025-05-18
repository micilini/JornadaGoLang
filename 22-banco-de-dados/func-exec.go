package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ExecQuery(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar query: %w", err)
	}
	return result, nil
}

func main() {
	// Conecta sem banco específico para criar o banco
	dbRoot, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer dbRoot.Close()

	_, err = ExecQuery(dbRoot, "CREATE DATABASE IF NOT EXISTS jornadagolang")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Banco de dados criado/verificado.")

	// Agora conecta no banco jornadagolang
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/jornadagolang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = ExecQuery(db, "DROP TABLE IF EXISTS users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = ExecQuery(db, `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			name VARCHAR(120),
			email VARCHAR(255),
			password VARCHAR(60),
			PRIMARY KEY (id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tabela users criada com sucesso!")
}
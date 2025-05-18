tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}

_, err = tx.Exec("INSERT INTO usuarios (nome, email) VALUES (?, ?)", "João", "joao@ex.com")
if err != nil {
	tx.Rollback()
	log.Fatal("Erro ao inserir:", err)
}

_, err = tx.Exec("DELETE FROM perfis WHERE usuario_id = ?", 5)
if err != nil {
	tx.Rollback()
	log.Fatal("Erro ao deletar perfil:", err)
}

err = tx.Commit()
if err != nil {
	log.Fatal("Erro ao confirmar transação:", err)
}
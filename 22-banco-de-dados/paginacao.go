func ListarUsuariosPaginado(db *sql.DB, page, pageSize int) ([]Usuario, error) {
	offset := (page - 1) * pageSize
	rows, err := db.Query("SELECT id, nome, email, criado_em FROM usuarios ORDER BY id LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome, &u.Email, &u.CriadoEm)
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}
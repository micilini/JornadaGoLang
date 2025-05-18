type Usuario struct {
	ID        int
	Nome      string
	Email     string
	CriadoEm  time.Time
}

func GetUsuarioByID(db *sql.DB, id int) (*Usuario, error) {
	row := db.QueryRow("SELECT id, nome, email, criado_em FROM usuarios WHERE id = ?", id)

	var u Usuario
	err := row.Scan(&u.ID, &u.Nome, &u.Email, &u.CriadoEm)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
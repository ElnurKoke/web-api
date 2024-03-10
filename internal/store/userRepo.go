package store

func (s *Store) UpdateUserName(id int, username string) error {
	query := `UPDATE user SET username = $1 WHERE id= $2;`
	if _, err := s.db.Exec(query, username, id); err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateUserEmail(id int, email string) error {
	query := `UPDATE user SET email = $1 WHERE id= $2;`
	if _, err := s.db.Exec(query, email, id); err != nil {
		return err
	}
	return nil
}

func (s *Store) CheckUserByName(username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE username = ?) AS UE_exists;"
	row := s.db.QueryRow(query, username)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *Store) CheckUserByEmail(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = ? ) AS UE_exists;"
	row := s.db.QueryRow(query, email)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

package database

func (d *dbConnector) AddToken(userid uint, publicKey string) error {
	stmt := `SELECT id from employees where id=$1`
	_, err := d.db.Exec(stmt, userid)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO securitytokens(user_id, publickey) values($1, $2)`
	_, err = d.db.Exec(stmt, userid, publicKey)
	if err != nil {
		return err
	}
	return nil
}
